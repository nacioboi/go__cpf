package definitions

import (
	"fmt"
)

func _inner_LogImplementation(
	requested_level int,
	current_level int,
	handlers map[int]CustomOutputHandlerEntry,
	msg string,
) {
	if requested_level <= current_level {
		for _, handler := range handlers {
			i := handler.I
			if i == 0 { // custom func
				handler.H(msg)
			} else if i == 1 { // *os.File
				n, err := handler.F.Write([]byte(msg))
				if err != nil {
					panic(err)
				}
				if n < len(msg) {
					panic("not all bytes were written")
				}
			} else {
				panic("something horrible happened")
			}
		}
	}
}

func LogImplementation(
	requested_level int,
	current_level int,
	handlers map[int]CustomOutputHandlerEntry,
	prefix_handler func() string,
	interval_amount *uint16,
	interval_count_downs map[string]uint16,
	format string,
	args ...interface{},
) {
	msg := fmt.Sprintf(
		"%s%s",
		prefix_handler(),
		fmt.Sprintf(format, args...),
	)
	if *interval_amount != 0 {
		if _, inside := interval_count_downs[format]; !inside {
			interval_count_downs[format] = *interval_amount
		}
		if interval_count_downs[format] == *interval_amount {
			interval_count_downs[format] = 0
			_inner_LogImplementation(
				requested_level,
				current_level,
				handlers,
				msg,
			)
		} else {
			interval_count_downs[format]++
		}
	} else {
		_inner_LogImplementation(
			requested_level,
			current_level,
			handlers,
			msg,
		)
	}
}

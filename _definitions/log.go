package definitions

import (
	"fmt"
)

func _inner_LogImplementation(
	requested_level int,
	current_level int,
	handlers map[int]CustomOutputHandlerEntry,
	prefix_handler func() string,
	interval_amount *uint16,
	interval_count_down *uint16,
	format string,
	args ...interface{},
) {
	if requested_level <= current_level {
		msg := fmt.Sprintf(
			"%s%s",
			prefix_handler(),
			fmt.Sprintf(format, args...),
		)
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
			*interval_count_down = *interval_amount
		}
	}
}

func LogImplementation(
	requested_level int,
	current_level int,
	handlers map[int]CustomOutputHandlerEntry,
	prefix_handler func() string,
	interval_amount *uint16,
	interval_count_down *uint16,
	format string,
	args ...interface{},
) {
	if *interval_amount != 0 {
		if *interval_count_down == 0 {
			_inner_LogImplementation(
				requested_level,
				current_level,
				handlers,
				prefix_handler,
				interval_amount,
				interval_count_down,
				format,
				args...,
			)
		} else {
			(*interval_count_down)--
		}
	} else {
		_inner_LogImplementation(
			requested_level,
			current_level,
			handlers,
			prefix_handler,
			interval_amount,
			interval_count_down,
			format,
			args...,
		)
	}
}

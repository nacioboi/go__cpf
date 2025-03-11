package definitions

import "fmt"

func LogImplementation(
	level int,
	output_func func(int, string),
	prefix_handler func() string,
	interval_amount *uint16,
	interval_count_down *uint16,
	format string,
	args ...interface{},
) {
	if *interval_amount != 0 {
		if *interval_count_down == 0 {
			output_func(
				level,
				fmt.Sprintf(
					"%s%s",
					prefix_handler(),
					fmt.Sprintf(format, args...),
				),
			)
			*interval_count_down = *interval_amount
		} else {
			(*interval_count_down)--
		}
	} else {
		output_func(
			level,
			fmt.Sprintf(
				"%s%s",
				prefix_handler(),
				fmt.Sprintf(format, args...),
			),
		)
	}
}

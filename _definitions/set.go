package definitions

import (
	"fmt"
	"io"

	"github.com/nacioboi/go_cpf/cpf_options"
)

func SetImplementation(
	key cpf_options.T__Option,
	value interface{},
	interval_amount *uint16,
	interval_count_down *uint16,
	c__MAX_INTERVAL_AMOUNT uint16,
	output_func *func(int, string),
	prefix_handler *func() string,
	log_level *int,
) {
	switch key {
	case cpf_options.PRINT_IN_INTERVALS:
		switch v := value.(type) {
		case int:
			if v < 0 {
				panic("cpf.Set: value is negative")
			}
			if v > int(c__MAX_INTERVAL_AMOUNT) {
				panic("cpf.Set: value is too large")
			}
			*interval_amount = uint16(v)
			*interval_count_down = *interval_amount
		default:
			panic("cpf.Set: value is not int")
		}
	case cpf_options.OUTPUT_HANDLER:
		switch v := value.(type) {
		case func(int, string):
			*output_func = v
		case io.Writer:
			*output_func = func(l int, s string) { fmt.Fprint(v, s) }
		default:
			panic("cpf.Set: value must be func(int, string) or io.Writer")
		}
	case cpf_options.PREFIX_HANDLER:
		switch v := value.(type) {
		case func() string:
			*prefix_handler = v
		default:
			panic("cpf.Set: value must be func() string")
		}
	case cpf_options.LOG_LEVEL:
		switch v := value.(type) {
		case int:
			if v < 0 {
				panic("cpf.Set: value is negative")
			}
			if v > 255 {
				panic("cpf.Set: value > 255")
			}
			*log_level = v
		default:
			panic("cpf.Set: value is not int")
		}
	}
}

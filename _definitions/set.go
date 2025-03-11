package definitions

import (
	"github.com/nacioboi/go__cpf/cpf_options"
)

func SetImplementation(
	key cpf_options.T__Option,
	value interface{},
	interval_amount *uint16,
	interval_count_down *uint16,
	c__MAX_INTERVAL_AMOUNT uint16,
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

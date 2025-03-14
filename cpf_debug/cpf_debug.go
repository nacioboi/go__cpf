package cpf_debug

import (
	definitions "github.com/nacioboi/go__cpf/_definitions"
	"github.com/nacioboi/go__cpf/cpf_options"
)

const c__MAX_INTERVAL_AMOUNT = 1<<16 - 1

const DEFAULT_HANDLER_ID = 1<<32 - 1

var (
	interval_count_downs = make(map[string]uint16)
	interval_amount      uint16

	log_level int

	output_handlers = make(map[int]definitions.CustomOutputHandlerEntry, 0)

	prefix_handler func() string

	default_output_handler = func(message string) {
		print(message)
	}
)

func init() {
	entry := definitions.CustomOutputHandlerEntry{
		I: 0,
		H: default_output_handler,
	}
	output_handlers[1<<32-1] = entry
}

func Add(key int, value interface{}) {
	definitions.AddImplementation(
		key,
		value,
		output_handlers,
		default_output_handler,
	)
}

func Del(key int) {
	definitions.DelImplementation(
		key,
		output_handlers,
	)
}

func Set(key cpf_options.T__Option, value interface{}) {
	definitions.SetImplementation(
		key,
		value,
		&interval_amount,
		c__MAX_INTERVAL_AMOUNT,
		&prefix_handler,
		&log_level,
	)
}

func Log(level int, format string, args ...interface{}) {
	definitions.LogImplementation(
		level,
		log_level,
		output_handlers,
		prefix_handler,
		&interval_amount,
		interval_count_downs,
		format,
		args...,
	)
}

func Formatted(out *string, format string, args ...interface{}) {
	definitions.FormattedImplementation(out, format, args...)
}

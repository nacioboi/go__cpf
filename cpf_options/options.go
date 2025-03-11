package cpf_options

type T__Option uint8

const (
	PRINT_IN_INTERVALS T__Option = 1 << iota
	OUTPUT_HANDLER
	LOG_LEVEL
	PREFIX_HANDLER
)

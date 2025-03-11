package cpf_options

type T__Option uint8

const (
	PRINT_IN_INTERVALS T__Option = 1 << iota
	LOG_LEVEL
	PREFIX_HANDLER
)

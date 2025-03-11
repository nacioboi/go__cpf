package definitions

import "os"

type CustomOutputHandler = func(string)

type CustomOutputHandlerEntry struct {
	I int
	H CustomOutputHandler
	F *os.File
}

func AddImplementation(
	key int,
	value interface{},
	handlers map[int]CustomOutputHandlerEntry,
	default_handler CustomOutputHandler,
) {
	if key == 1<<32-1 {
		if _, is_inside := handlers[key]; is_inside {
			panic("cpf.Add: key already within map")
		} else {
			entry := CustomOutputHandlerEntry{
				I: 0,
				H: default_handler,
			}
			handlers[key] = entry
			return
		}
	}

	if _, is_inside := handlers[key]; is_inside {
		panic("cpf.Add: key already exists")
	}
	switch value := value.(type) {
	case CustomOutputHandler:
		entry := CustomOutputHandlerEntry{
			I: 0,
			H: value,
		}
		handlers[key] = entry
	case *os.File:
		entry := CustomOutputHandlerEntry{
			I: 1,
			F: value,
		}
		handlers[key] = entry
	default:
		panic("cpf.Add: value must be `CustomOutputHandler` or `*os.File`")
	}
}

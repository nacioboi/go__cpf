package definitions

func DelImplementation(
	key int,
	handlers map[int]CustomOutputHandlerEntry,
) {
	if _, is_inside := handlers[key]; !is_inside {
		panic("cpf.Del: key not in map")
	}
	delete(handlers, key)
}

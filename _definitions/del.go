package definitions

func DelImplementation(
	key int,
	handlers map[int]CustomOutputHandlerEntry,
) {
	delete(handlers, key)
}

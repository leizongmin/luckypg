package nostd

func assert(ok bool) {
	if !ok {
		panic("assert failed")
	}
}

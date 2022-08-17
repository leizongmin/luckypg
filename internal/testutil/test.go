package testutil

func Assert(ok bool) {
	if !ok {
		panic("assert failed")
	}
}

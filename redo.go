package repeat

func RedoUntilOKFunc(num uint, fn func()bool) func() {
	return func() {
		RedoUntilOK(num, fn)
	}
}

// RedoUntilOK will call 'fn' until it returns 'true', but will only do this at most 'num' timnes.
func RedoUntilOK(num uint, fn func()bool) {
	if nil == fn {
		panic(errNilFunc)
	}

	for i := uint(0); i<num; i++ {
		ok := fn()
		if ok {
			break
		}
	}
}

package headacherr

func Defer(f func(err error)) {
	if r := recover(); r == nil {
		return
	} else if err, ok := r.(error); !ok {
		panic(err)
	} else if f != nil {
		f(err)
	}
}

func Try[T any](result T, err error) T {
	if err == nil {
		return result
	}

	panic(err)
}

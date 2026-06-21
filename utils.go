package err

func Safely(f func(), onPanic func(v any)) {
	defer func() {
		if r := recover(); r != nil {
			onPanic(r)
		}
	}()

	f()
}
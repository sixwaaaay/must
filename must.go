package must

// the defer statement is executed even if the function panics
// so in the main() function,defer and Must() can be used to simplify the error handling

// do not use in library code, only in top level main() functions where you want to avoid error handling boilerplate

// Must panics if err is not nil.
// use this to avoid if err != nil { panic(err) } boilerplate in main()
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// RunE panics if err is not nil.
// use this to avoid if err != nil { panic(err) } boilerplate in main()
func RunE(err error) {
	if err != nil {
		panic(err)
	}
}

package halpers

func IfPanicError(err error) {
	if err != nil {
		panic(err)
	}

}

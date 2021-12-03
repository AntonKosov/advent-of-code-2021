package aoc

func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

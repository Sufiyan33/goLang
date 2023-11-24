package main

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

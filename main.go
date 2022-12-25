package main

import (
	"os"

	"github.com/temphia/temphia/code/tools/repobuild"
)

func main() {

	fout, err := os.ReadFile("repo.json")
	handleErr(err)

	rb, err := repobuild.New(fout)
	handleErr(err)

	result, err := rb.BuildAll()
	handleErr(err)

	println(result)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

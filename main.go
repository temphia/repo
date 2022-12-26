package main

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/tools/repobuild"
)

func main() {

	fout, err := os.ReadFile("repo.json")
	handleErr(err)

	rb, err := repobuild.New(fout)
	handleErr(err)

	result, err := rb.BuildAll()
	handleErr(err)

	pp.Println(result)

	for _, err2 := range result.ErroredItems {
		pp.Println(err2.Error())

	}

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
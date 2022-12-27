package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/repo/code"
)

func main() {

	fout, err := os.ReadFile("repo.json")
	handleErr(err)

	rb, err := code.New(fout)
	handleErr(err)

	err = rb.BuildAll()
	handleErr(err)

	fmt.Printf("Out of %d, %d built sucessfully and %d errored out \n", len(rb.Config.Items), len(rb.Outputs), len(rb.ErroredItems))

	for k, err2 := range rb.ErroredItems {
		fmt.Printf(" [ %s ] errored |> %+v ", k, err2)
		fmt.Println(err2.Error())
	}

	for k, outFolder := range rb.Outputs {
		fmt.Printf(" [ %s ] output |> %+v ", k, outFolder)

	}

	pp.Println(rb.IndexAll(false))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

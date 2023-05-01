package main

import "github.com/k0kubun/pp"

func main() {

	pp.Println()

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"log"

	"github.com/k0kubun/pp"
	rcmd "github.com/temphia/temphia/code/tools/repobuild/cmd"
)

func main() {
	err := rcmd.Run(".repo.yaml")
	if err != nil {
		pp.Println(err)
		log.Println(err.Error())
	}
}

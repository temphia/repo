package main

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	rcmd "github.com/temphia/temphia/code/tools/repobuild/cmd"
)

func main() {

	err := xutils.CreateIfNotExits("build")
	if err != nil {
		panic(err)
	}

	pp.Println()

	err = rcmd.Run(".repo.yaml")
	if err != nil {
		pp.Println(err)
		panic(err.Error())
	}
}

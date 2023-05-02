package main

import (
	"github.com/k0kubun/pp"
	rcmd "github.com/temphia/temphia/code/tools/repobuild/cmd"
)

func main() {

	pp.Println(rcmd.Run(".repo.yaml"))

}

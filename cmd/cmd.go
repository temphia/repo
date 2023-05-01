package cmd

import (
	"encoding/json"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/repo/pkg/builder"
	"github.com/temphia/repo/pkg/models"
)

func Run() {

	cbytes, err := os.ReadFile("repo.json")
	if err != nil {
		panic(err)
	}

	conf := &models.BuildConfig{}
	err = json.Unmarshal(cbytes, conf)
	if err != nil {
		panic(err)
	}

	builder := builder.New(conf)

	pp.Println(builder.Build())

}

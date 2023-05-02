package cmd

import (
	"os"

	"github.com/temphia/repo/pkg/builder"
	"github.com/temphia/repo/pkg/models"
	"gopkg.in/yaml.v2"
)

type Options struct {
	RepoFile string
}

func Run(opts *Options) {

	cbytes, err := os.ReadFile(opts.RepoFile)
	if err != nil {
		panic(err)
	}

	conf := &models.BuildConfig{}
	err = yaml.Unmarshal(cbytes, conf)
	if err != nil {
		panic(err)
	}

	builder := builder.New(conf)

	err = builder.Build()
	if err != nil {
		panic(err)
	}

	builder.PrintResult()
}

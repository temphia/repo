package builder

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/k0kubun/pp"
	"github.com/temphia/repo/pkg/index"
	"github.com/temphia/repo/pkg/models"
)

type RepoBuilder struct {
	config *models.BuildConfig

	// build stage states
	ErroredItems map[string]error
	Outputs      map[string]string

	// index stage states
	indexer *index.Indexer

	repoCache map[string]*git.Repository
}

func New(conf *models.BuildConfig) *RepoBuilder {

	return &RepoBuilder{
		config:       conf,
		indexer:      index.New(path.Join(conf.OutputFolder, "db.json")),
		ErroredItems: make(map[string]error),
		Outputs:      make(map[string]string),
		repoCache:    make(map[string]*git.Repository),
	}
}

func (rb *RepoBuilder) Build() error {

	os.RemoveAll(rb.config.BuildFolder)

	for k := range rb.config.Items {

		ofolder, err := rb.buildItem(k)
		if err != nil {
			rb.ErroredItems[k] = err
			continue
		}
		rb.Outputs[k] = ofolder
	}

	return nil
}

func (rb *RepoBuilder) PrintResult() {
	for k, err := range rb.ErroredItems {
		pp.Println("@err", k, err)
	}

	for k, v := range rb.Outputs {
		pp.Println("@build_ok", k, v)
	}

}

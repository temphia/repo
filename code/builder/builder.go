package builder

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/repo/code/models"
)

type RepoBuilder struct {
	config *models.BuildConfig
}

func New(conf *models.BuildConfig) *RepoBuilder {
	return &RepoBuilder{
		config: conf,
	}
}

func (rb *RepoBuilder) Build() error {

	for _, ri := range rb.config.Items {
		pp.Println(ri)
	}

	return nil
}

func (rb *RepoBuilder) buildItem() error {

	return nil
}

/*


type RepoBuild struct {
	Config *BuildConfig

	// build stage states
	ErroredItems map[string]error
	Outputs      map[string]string

	// index stage states
	db *DB
}

func New(conf []byte) (*RepoBuild, error) {
	bconf := &BuildConfig{}

	err := json.Unmarshal(conf, bconf)
	if err != nil {
		return nil, err
	}

	return &RepoBuild{
		Config:       bconf,
		ErroredItems: make(map[string]error),
		Outputs:      make(map[string]string),
		db:           nil,
	}, nil

}

func (rb *RepoBuild) BuildAll() error {

	for k := range rb.Config.Items {

		ofolder, err := rb.BuildOne(k, false)
		if err != nil {
			rb.ErroredItems[k] = err
			continue
		}
		rb.Outputs[k] = ofolder
	}

	return nil
}

func (rb *RepoBuild) BuildOne(name string, zip bool) (string, error) {
	of, err := rb.buildItem(name)
	if err != nil {
		return "", err
	}

	if !zip {
		return of, nil
	}

	panic("Zip not implemented")
}

func (rb *RepoBuild) IndexAll(ignoreOld bool) error {
	rb.initDB(ignoreOld)
	for k, v := range rb.Outputs {
		err := rb.indexItem(k, v)
		if err != nil {
			return err
		}
	}

	return rb.saveDB()
}


*/

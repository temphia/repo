package index

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/tidwall/gjson"
)

type Indexer struct {
	db   *DB
	file string
}

func New(file string) *Indexer {

	db := &DB{
		GroupIndex: make(map[string][]string),
		TagIndex:   make(map[string][]string),
		Items:      make(map[string]map[string]any),
	}

	fout, err := os.ReadFile(file)
	if err == nil {
		err := json.Unmarshal(fout, db)
		if err != nil {
			panic(err)
		}
	}

	return &Indexer{
		db:   db,
		file: file,
	}
}

func (dbi *Indexer) IndexItem(name, folder string) error {

	out, err := os.ReadFile(path.Join(folder, "index.json"))
	if err != nil {
		return err
	}

	slug := gjson.GetBytes(out, "slug").String()
	if slug != name {
		return fmt.Errorf("err: slug mismatch expected %s, got %s", name, slug)
	}

	gtype := gjson.GetBytes(out, "type").String()

	groups, ok := dbi.db.GroupIndex[gtype]
	if !ok {
		groups = []string{name}
	}

	dbi.db.GroupIndex[gtype] = groups

	data := make(map[string]any)

	err = json.Unmarshal(out, &data)
	if err != nil {
		return err
	}

	dbi.db.Items[name] = data

	// fixme => also index tags

	return nil
}

func (dbi *Indexer) Save() error {
	out, err := json.MarshalIndent(dbi.db, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbi.file, out, 0755)
}

type DB struct {
	GroupIndex map[string][]string       `json:"group_index" yaml:"group_index"`
	TagIndex   map[string][]string       `json:"tag_index" yaml:"tag_index"`
	Items      map[string]map[string]any `json:"items" yaml:"items"`
}

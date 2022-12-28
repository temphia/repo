package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/tidwall/gjson"
)

func (rb *RepoBuild) indexItem(name, folder string) error {

	out, err := os.ReadFile(path.Join(folder, "index.json"))
	if err != nil {
		return err
	}

	slug := gjson.GetBytes(out, "slug").String()
	if slug != name {
		return fmt.Errorf("err: slug mismatch expected %s, got %s", name, slug)
	}

	gtype := gjson.GetBytes(out, "type").String()

	groups, ok := rb.db.GroupIndex[gtype]
	if !ok {
		groups = []string{name}
	}

	rb.db.GroupIndex[gtype] = groups

	data := make(map[string]any)

	err = json.Unmarshal(out, &data)
	if err != nil {
		return err
	}

	rb.db.Items[name] = data

	// fixme => also index tags

	return nil
}

func (rb *RepoBuild) initDB(ignoreOld bool) {
	if rb.db == nil {
		rb.db = &DB{
			GroupIndex: make(map[string][]string),
			TagIndex:   make(map[string][]string),
			Items:      make(map[string]map[string]any),
		}

		if !ignoreOld {
			fout, err := os.ReadFile(path.Join(rb.Config.OutputFolder, "db.json"))
			if err == nil {
				err := json.Unmarshal(fout, rb.db)
				if err != nil {
					panic(err)
				}
			}
		}
	}

}

func (rb *RepoBuild) saveDB() error {

	out, err := json.Marshal(rb.db)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(rb.Config.OutputFolder, "db.json"), out, 0755)
}

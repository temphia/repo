package code

import (
	"encoding/json"
	"os"
	"path"
)

func (rb *RepoBuild) indexItem(name, folder string) error {

	out, err := os.ReadFile(path.Join(folder, "index.json"))
	if err != nil {
		return err
	}

	manifest := make(map[string]any)

	err = json.Unmarshal(out, &manifest)
	if err != nil {
		return err
	}

	gtype := manifest["group"].(string)

	groups, ok := rb.db.GroupIndex[gtype]
	if !ok {
		groups = []string{name}
	}

	rb.db.GroupIndex[gtype] = groups

	// fixme => also index tags

	return nil
}

func (rb *RepoBuild) initDB(ignoreOld bool) {
	if rb.db == nil {
		rb.db = &DB{
			GroupIndex: make(map[string][]string),
			TagIndex:   make(map[string][]string),
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

package builder

import (
	"fmt"
	"os"
	"path"

	"github.com/temphia/repo/code/utils"
	"github.com/tidwall/gjson"
)

func copyBprintFiles(artifactFolder, outputFolder string) error {

	out, err := os.ReadFile(path.Join(artifactFolder, "index.json"))
	if err != nil {
		return err
	}

	err = utils.CreateIfNotExists(outputFolder, 0755)
	if err != nil {
		return err
	}

	result := gjson.GetBytes(out, "files").Array()
	for _, r := range result {
		fmt.Println("@copying file ", r)
		file := r.String()
		err := utils.Copy(
			path.Join(artifactFolder, file),
			path.Join(outputFolder, file),
		)
		if err != nil {
			return err
		}
	}

	return utils.Copy(
		path.Join(artifactFolder, "index.json"),
		path.Join(outputFolder, "index.json"),
	)

}

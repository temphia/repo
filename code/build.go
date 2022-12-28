package code

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/libx/xutils"
	"github.com/tidwall/gjson"
)

func (rb *RepoBuild) buildItem(name string) (string, error) {

	fmt.Println("Building ", name)

	item := rb.Config.Items[name]

	buildPath := path.Join(rb.Config.BuildFolder, name)
	outputPath := path.Join(rb.Config.OutputFolder, name)

	err := xutils.CreateIfNotExits(buildPath)
	if err != nil {
		return "", err
	}

	_, err = git.PlainClone(buildPath, false, &git.CloneOptions{
		URL:           item.GitURL,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(item.Branch),
		SingleBranch:  true,
		Depth:         1,
	})

	if err != nil {
		if !errors.Is(git.ErrRepositoryAlreadyExists, err) {
			return "", err
		}
	}

	curr, _ := os.Getwd()
	os.Chdir(path.Join(curr, buildPath))
	cmd := exec.Command(item.BuildCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	os.Chdir(curr)
	if err != nil {
		return "", err
	}

	artifactFolder := path.Join(buildPath, item.OutputFolder)

	pp.Println("@copying_form", artifactFolder, "->", outputPath)

	return outputPath, copyBprintFiles(artifactFolder, outputPath)
}

func copyBprintFiles(artifactFolder, outputFolder string) error {

	out, err := os.ReadFile(path.Join(artifactFolder, "index.json"))
	if err != nil {
		return err
	}

	err = CreateIfNotExists(outputFolder, 0755)
	if err != nil {
		return err
	}

	result := gjson.GetBytes(out, "files").Array()
	for _, r := range result {
		fmt.Println("@copying file ", r)
		file := r.String()
		err := Copy(
			path.Join(artifactFolder, file),
			path.Join(outputFolder, file),
		)
		if err != nil {
			return err
		}
	}

	return Copy(
		path.Join(artifactFolder, "index.json"),
		path.Join(outputFolder, "index.json"),
	)

}

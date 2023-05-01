package builder

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/k0kubun/pp"
	"github.com/temphia/repo/pkg/utils"
)

func (rb *RepoBuilder) buildItem(name string) (string, error) {

	fmt.Println("Building ", name)

	item := rb.config.Items[name]

	buildPath := rb.hashedBuidlPath(item.GitURL)
	outputPath := path.Join(rb.config.OutputFolder, name)

	err := utils.CreateIfNotExists(buildPath, 0755)
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
	cmd := exec.Command(item.BuildCMD)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	os.Chdir(curr)
	if err != nil {
		return "", err
	}

	artifactFolder := path.Join(buildPath, item.Output)

	pp.Println("@copying_form", artifactFolder, "->", outputPath)

	return outputPath, copyBprintFiles(artifactFolder, outputPath)

}

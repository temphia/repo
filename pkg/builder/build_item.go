package builder

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/temphia/repo/pkg/utils"
)

func (rb *RepoBuilder) buildItem(name string) (string, error) {

	fmt.Println("Building ", name)

	item := rb.config.Items[name]

	buildPath := rb.hashedBuidlPath(item.GitURL)

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

	err = rb.runBuild(buildPath, item.BuildCMD)
	if err != nil {
		panic(err)
	}

	// 	artifactFolder := path.Join(buildPath, item.Output)

	// outputPath := path.Join(rb.config.OutputFolder, name)

	// pp.Println("@copying_form", artifactFolder, "->", outputPath)

	//return outputPath, copyBprintFiles(artifactFolder, outputPath)

	return "", nil

}

func (rb *RepoBuilder) runBuild(workFolder, buildcmd string) error {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	vol := fmt.Sprintf("%s/%s:/work", wd, workFolder)

	cmd := exec.Command(
		"docker",
		"run",
		"-it",
		"-v",
		vol,
		"ghcr.io/temphia/temphia_buildpack:latest",
		buildcmd,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

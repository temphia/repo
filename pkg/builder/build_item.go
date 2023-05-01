package builder

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/temphia/repo/pkg/utils"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/tools/bdev"
	"gopkg.in/yaml.v2"
)

func (rb *RepoBuilder) buildItem(name string) (string, error) {

	fmt.Println("Building ", name)

	item := rb.config.Items[name]

	buildPath := rb.hashedBuidlPath(item.GitURL)

	err := utils.CreateIfNotExists(buildPath, 0755)
	if err != nil {
		return "", err
	}

	repo, err := git.PlainClone(buildPath, false, &git.CloneOptions{
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

	if repo != nil {
		rb.repoCache[buildPath] = repo
	} else {
		repo = rb.repoCache[buildPath]
	}

	headRef, err := repo.Head()
	if err != nil {
		panic(err)
	}

	versionHash := headRef.String()[:7]

	err = rb.runBuild(buildPath, item.BuildCMD)
	if err != nil {
		panic(err)
	}

	err = rb.copyArtifact(buildPath, name, item.BprintFile, versionHash)
	if err != nil {
		return "", err
	}

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
		"--rm",
		"-v",
		vol,
		"ghcr.io/temphia/temphia_buildpack:latest",
		buildcmd,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func (rb *RepoBuilder) copyArtifact(basePath, name, bprintFile, version string) error {
	out, err := os.ReadFile(path.Join(basePath, bprintFile))
	if err != nil {
		return err
	}

	lbprint := &xbprint.LocalBprint{}
	err = yaml.Unmarshal(out, lbprint)
	if err != nil {
		return err
	}
	return bdev.ZipIt(lbprint, path.Join(rb.config.OutputFolder, fmt.Sprintf("%s_%s.zip", name, version)))
}

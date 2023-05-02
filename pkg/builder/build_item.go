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
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/tools/bdev"
	"gopkg.in/yaml.v2"
)

func (rb *RepoBuilder) buildItem(name string) (string, error) {

	fmt.Println("Building ", name)

	item := rb.config.Items[name]

	buildPath := rb.hashedBuidlPath(item.GitURL)

	// clone repo
	versionHash, err := rb.gitClone(buildPath, item.GitURL, item.Branch)
	if err != nil {
		return "", err
	}

	// actual build
	err = rb.runBuild(buildPath, item.BuildCMD)
	if err != nil {
		panic(err)
	}

	// copy artifacts
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

func (rb *RepoBuilder) gitClone(path, url, branch string) (string, error) {
	err := utils.CreateIfNotExists(path, 0755)
	if err != nil {
		return "", err
	}

	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:           url,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		SingleBranch:  true,
		Depth:         1,
	})

	if err != nil {
		if !errors.Is(git.ErrRepositoryAlreadyExists, err) {
			return "", err
		}
	}

	if repo != nil {
		rb.repoCache[path] = repo
	} else {
		repo = rb.repoCache[path]
	}

	headRef, err := repo.Head()
	if err != nil {
		panic(err)
	}

	return headRef.String()[:7], nil
}

func (rb *RepoBuilder) copyArtifact(basePath, name, bprintFile, version string) error {
	out, err := os.ReadFile(path.Join(basePath, bprintFile))
	if err != nil {
		return err
	}

	pp.Println(string(out))

	lbprint := &xbprint.LocalBprint{}
	err = yaml.Unmarshal(out, lbprint)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s_%s.zip", name, version)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(basePath)
	if err != nil {
		panic(err)
	}

	defer func() {
		os.Remove(filename)

	}()

	err = bdev.ZipIt(lbprint, filename)
	os.Chdir(wd)

	if err != nil {
		return err
	}

	distpath := path.Join(rb.config.OutputFolder, name)

	utils.CreateIfNotExists(distpath, 0755)

	err = utils.Copy(
		path.Join(basePath, filename),
		path.Join(distpath, fmt.Sprintf("%s.zip", version)),
	)
	if err != nil {
		pp.Println(err.Error())
		return err
	}

	return nil
}

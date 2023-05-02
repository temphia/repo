package main

import "github.com/temphia/repo/cmd"

func main() {
	cmd.Run(&cmd.Options{
		RepoFile: "repo.json",
	})
}

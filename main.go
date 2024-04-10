package main

import (
	//"fmt"
	//"os"

	"github.com/lukasnxyz/gitvis/src/user"
	//"github.com/lukasnxyz/gitvis/src/draw"
	//"github.com/lukasnxyz/gitvis/src/localdata"

	"github.com/google/go-github/v61/github"
)

func main() {
	// handle command line input
	// check local repo if .git by default and siplay commits for repo
	// flag '-u' to input github username
	// flag '-l' to only check commits to local repos on system
	// flag '-h' for help

	client := github.NewClient(nil)
	gitubRepos := user.FetchUserRepos(client, "lukasnxyz")
	repos := user.ConvUserRepos(gitubRepos)

	user.PrintUserRepos(repos)
}

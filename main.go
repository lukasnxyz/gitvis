package main

import (
	"fmt"
	"os"

	//"github.com/lukasnxyz/gitvis/src/user"
	"github.com/lukasnxyz/gitvis/src/localrepo"
	//"github.com/google/go-github/v61/github"
)

func main() {
	// handle command line input
	// check local repo if .git by default and display commits for repo
	// flag '-u' to input github username
	// flag '-h' for help

	//client := github.NewClient(nil)
	//gitubRepos := user.FetchUserRepos(client, "lukasnxyz")
	//repos := user.ConvUserRepos(gitubRepos)
	//user.PrintUserRepos(repos)

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	repo := localrepo.Fet(path)
	fmt.Println(repo)

	for i := 0; i < repo.NumOfCommits; i++ {
		fmt.Println(repo.Commits[i])
	}
}

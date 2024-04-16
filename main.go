package main

import (
	"flag"
	//"fmt"
	"os"

	"github.com/lukasnxyz/gitvis/src/display"
	"github.com/lukasnxyz/gitvis/src/localrepo"
)

/*
	notes
	- show commit history like on GitHub by week rows
	- be able to inspect a commit
	- a less feature for the output?
*/

func main() {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var path string
	localRepoFlag := flag.Bool("l", false, "visualize commits of git repository in current directory")
	flag.StringVar(&path, "d", currPath, "directory for visualization of local git repository")
	ghRepoFlag := flag.Bool("g", false, "visualize commits of git repository in on GitHub")
	ghUserFlag := flag.Bool("u", false, "visualize commits of user on GitHub")

	flag.Parse()

	// default to local, if not, don't pass local
	if *localRepoFlag {
		repo := localrepo.Fet(path)
		// repo.Commits is ordered first commits last
		// repo.LinesOfCode is also not accurate, just counts all text lines in repo

		display.Visualize(repo)
	}

	if *ghRepoFlag {
	}

	if *ghUserFlag {
	}
}

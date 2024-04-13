package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lukasnxyz/gitvis/src/localrepo"
)

// Notes
// - cli input
// - don't use panic for errors

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

	if *localRepoFlag {
		repo := localrepo.Fet(path)
		fmt.Println(repo)
		fmt.Println()

		for i := 0; i < repo.NumOfCommits; i++ {
			fmt.Println(repo.Commits[i])
		}

		//display.Vis(repo)
	}

	if *ghRepoFlag {
	}

	if *ghUserFlag {
	}
}

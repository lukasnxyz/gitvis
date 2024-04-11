package main

import (
	"fmt"
	"os"

	"github.com/lukasnxyz/gitvis/src/localrepo"
)

func main() {
	// handle command line input

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	repo := localrepo.Fet(path)
	fmt.Println(repo)
	fmt.Println()

	for i := 0; i < repo.NumOfCommits; i++ {
		fmt.Println(repo.Commits[i])
	}

	//display.Vis(repo)
}

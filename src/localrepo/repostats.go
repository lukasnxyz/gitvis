package localrepo

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type Repo struct {
	Name string
}

func (r Repo) String() {
	return fmt.Sprintf()
}

func Stats() {
	commits := process()
	printCommits(commits)
}

func fillCommits() map[int]int {
	path := "."
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	iter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	offset := calcOffset()
	err = iter.ForEach(func(c *object.Commit) error {
	}

}

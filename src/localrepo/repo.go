package localrepo

import (
	"fmt"
	//"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const (
	daysLastSixMonths = 183
	outOfRange        = 99999
)

type commit struct {
	Time    time.Time
	Message      string
	Author       string
	Id           plumbing.Hash
	LinesAdded   int
	LinesDeleted int
}

type Repository struct {
	Name         string
	Path         string
	Status       string
	Commits      []commit
	NumOfCommits int
	Lines        int
}

func (r Repository) String() string {
	return fmt.Sprintf("Repo [\n\tName: %s\n\tPath: %s\n\tStatus: %s\n\tNumber of commits: %d\n\tLines of code: %d\n]",
		r.Name, r.Path, r.Status, r.NumOfCommits, r.Lines)
}

func (c commit) String() string {
	return fmt.Sprintf("Commit [\n\tTime: %s\n\tMessage: %s\tAuthor: %s\n\tId: %s\n\tDiff: -%d	+%d\n]",
		c.Time.String(), c.Message, c.Author, c.Id, c.LinesDeleted, c.LinesAdded)
}

func Fet(path string) Repository {
	repo := NewRepository(path)
	repo.fillCommits()

	return repo
}

func NewRepository(path string) (repo Repository) {
	repo.Name = ""
	repo.Path = path
	repo.Status = ""
	//repo.Commits
	repo.NumOfCommits = 0
	repo.Lines = 0

	return
}

func (r *Repository) fillCommits() {
	repo, err := git.PlainOpen(r.Path)
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

	iter.ForEach(func(c *object.Commit) error {
		nCommit := commit{
			Time:         c.Author.When,
			Message:      c.Message,
			Author:       c.Author.String(),
			Id:           c.ID(),
			LinesAdded:   0,
			LinesDeleted: 0,
		}

		stats, err := c.Stats()
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(stats); i++ {
			nCommit.LinesAdded += stats[i].Addition
			nCommit.LinesDeleted += stats[i].Deletion
		}

		r.NumOfCommits++ // this doesn't work
		r.Commits = append(r.Commits, nCommit)
		r.Lines += nCommit.LinesAdded - nCommit.LinesDeleted

		return nil
	})
}

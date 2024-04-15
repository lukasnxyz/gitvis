package localrepo

// back-end

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const (
	daysLastSixMonths = 183
	outOfRange        = 99999
)

type Repository struct {
	Name         string
	Path         string
	Status       string
	Commits      []Commit
	NumOfCommits int
	Lines        int
}

type Commit struct {
	Time         time.Time
	Message      string
	Author       string
	Email        string
	Id           plumbing.Hash
	LinesAdded   int
	LinesDeleted int
	Lines        int
}

func (r Repository) String() string {
	return fmt.Sprintf("Repo [\n\tName: %s\n\tPath: %s\n\tStatus: %s\n\tNumber of commits: %d\n\tLines of code: %d\n]",
		r.Name, r.Path, r.Status, r.NumOfCommits, r.Lines)
}

func (c Commit) String() string {
	return fmt.Sprintf("Commit [\n\tTime: %s\n\tMessage: %s\tAuthor: %s\n\tEmail: %s\n\tId: %s\n\tDiff: -%d	+%d\n\tLines: %d\n]",
		c.Time, c.Message, c.Author, c.Email, c.Id, c.LinesDeleted, c.LinesAdded, c.Lines)
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
	repo.NumOfCommits = 0
	repo.Lines = 0

	return
}

// break this up
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
		nCommit := Commit{
			Time:         c.Author.When,
			Message:      c.Message,
			Author:       c.Author.Name,
			Email:        c.Author.Email,
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

		r.NumOfCommits++
		r.Lines += nCommit.LinesAdded - nCommit.LinesDeleted

		r.Commits = append(r.Commits, nCommit)

		return nil
	})
}

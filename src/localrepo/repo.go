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
	Time         time.Time
	Message      string
	Author       string
	AuthorEmail  string
	CommitId     plumbing.Hash
	LinesAdded   int
	LinesDeleted int
}

type Repository struct {
	Name         string
	Path         string
	Status       string
	Commits      []commit
	NumOfCommits int
}

func (r Repository) String() string {
	return fmt.Sprintf("Repo [\n\tName: %s\n\tPath: %s\n\tStatus: %s\n\tNumber of commits: %d\n]",
		r.Name, r.Path, r.Status, r.NumOfCommits)
}

func (c commit) String() string {
	return fmt.Sprintf("Commit [\n\tTime: %s\n\tMessage: %s\n\tAuthor: %s\n\tEmail: %s\n\tCommitId: %s\n\tDiff: -%d	+%d\n]",
		c.Time.String(), c.Message, c.Author, c.AuthorEmail, c.CommitId, c.LinesDeleted, c.LinesAdded)
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

	// iter is an interface

	iter.ForEach(func(o *object.Commit) error {
		nCommit := commit{
			Time:        time.Now(),
			Message:     "",
			Author:      "",
			AuthorEmail: "",
			CommitId:    o.ID(),
			LinesAdded: 0,
			LinesDeleted: 0,
		}

		stats, err := o.Stats()
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(stats); i++ {
			nCommit.LinesAdded += stats[i].Addition
			nCommit.LinesDeleted += stats[i].Deletion
		}

		//fmt.Println(o.String())

		r.NumOfCommits++ // this doesn't work
		r.Commits = append(r.Commits, nCommit)

		return nil
	})
}

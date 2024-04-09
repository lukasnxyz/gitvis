package fetdata

import (
	"fmt"
	"log"
	"context"
	"os"

	"github.com/google/go-github/v61/github"
)

type UserRepo struct {
	RepoName string
	HtmlUrl  string
}

// Todo
// - parse the json output user repo list into custom slice of structs
// - write tests of course

func FetchRepos(user string) {
	client := github.NewClient(nil)
	ctx := context.Background()

	// list public repositories for org "github"
	opt := &github.RepositoryListByUserOptions{Type: "all", Sort: "updated"}
	repos, _, err := client.Repositories.ListByUser(ctx, user, opt)
	if err != nil {
		log.Fatal("could not reach: ", user, " ", err)
	}

	//fmt.Println(reflect.TypeOf(repos))
	//fmt.Println(repos)
	err = writeRepos("data.json", repos)
}

func writeRepos(filepath string, repos []*github.Repository) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, repo := range repos {
		fmt.Fprintln(f, repo)
	}

	log.Printf("write to %s was successful\n", filepath)

	return nil
}

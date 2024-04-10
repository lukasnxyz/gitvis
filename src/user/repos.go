package fetdata

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v61/github"
)

type UserRepo struct {
	Name        string
	HTMLURL         string
	CloneURL string
	Description string
}

// Todo
// - parse the json output user repo list into custom slice of structs
// - write tests of course

func FetchUserRepos(client *github.Client, user string) []*github.Repository {
	ctx := context.Background()

	// list public repositories for org "github"
	opt := &github.RepositoryListByUserOptions{Type: "all", Sort: "updated"}
	fmt.Printf("Fetching all public repositories of user '%s'\n", user)
	repos, _, err := client.Repositories.ListByUser(ctx, user, opt)
	if err != nil {
		log.Fatal("could not reach: ", user, " ", err)
	}

	return repos
}

func ConvUserRepos(repos []*github.Repository) (userRepos []UserRepo) {
	for i := 0; i < len(repos); i++ {
		userRepos = append(userRepos, UserRepo{
			Name: repos[i].GetName(),
			HTMLURL: repos[i].GetHTMLURL(),
			CloneURL: repos[i].GetCloneURL(),
			Description: repos[i].GetDescription(),
		})
	}

	return
}

func PrintUserRepos(userRepos []UserRepo) {
	for i := 0; i < len(userRepos); i++ {
		fmt.Println(userRepos[i])
	}
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

func (u UserRepo) String() string {
	return fmt.Sprintf("UserRepo: [\n\tName: '%s',\n\tURL: '%s',\n\tClone: '%s',\n\tDescription: '%s'\n]",
		u.Name, u.HTMLURL, u.CloneURL, u.Description)
}

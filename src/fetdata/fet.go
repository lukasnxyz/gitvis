package fetdata

import (
	"fmt"

	"github.com/google/go-github/v61/github"
)

	client := github.NewClient(nil)

	orgs, _, err := client.Organizations.List(nil, "willnorris", nil)
	if err != nil {
		log.Fatal(err)
	}


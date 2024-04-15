package display

// Front-end

import (
	"fmt"

	"github.com/lukasnxyz/gitvis/src/localrepo"
)

type DisplayNode struct {
}

type Display struct {
	Length int // term cols
	Width  int // term rows
	Nodes  []DisplayNode
}

func sort(repo localrepo.Repository) Display {
	display := Display{
		Length: 10,
		Width:  10,
	}

	return display
}

func Visualize(repo localrepo.Repository) {
	display := sort(repo)

	fmt.Println(repo.Name)
	fmt.Println(display.Length)
}

package display

// front-end

import (
	"fmt"
	"time"

	"github.com/lukasnxyz/gitvis/src/localrepo"
	"github.com/fatih/color"
)

type Color string
const (
	darkGreen Color = "#006400"
	lightGreen Color = "#90ee90"
	green Color = "#00ff00"
)

type DisplayNode struct {
	Color string
	Time time.Time
}

type Display struct {
	Length int // term cols
	Height  int // term rows
	Nodes  []DisplayNode
}

func sort(repo localrepo.Repository) Display {
	display := Display{
		Length: 52,
		Height: 7,
	}

	return display
}

func Visualize(repo localrepo.Repository) {
	fmt.Print(repo.StringShort())
	for _, commit := range repo.Commits {
		linesChanged := commit.LinesAdded + commit.LinesDeleted

		if linesChanged > 30 {
			color.Blue(commit.StringShort())
		} else if linesChanged > 15 {
			color.Green(commit.StringShort())
		}
	}
}

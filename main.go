package main

import (
	"os"

	"github.com/lukasnxyz/gitvis/draw"
	"github.com/lukasnxyz/gitvis/fetdata"
)

func main() {
	draw.PrintMsg(os.Stdout, "ready to go")

	fetdata.FetchRepos("lukasnxyz")
}

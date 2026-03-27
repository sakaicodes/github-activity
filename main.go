package main

import (
	"os"

	"github.com/sakaicodes/github-activity/cmd"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: github-activity <github-username>")
		os.Exit(1)
	}
	cmd.FilterEvent()
}

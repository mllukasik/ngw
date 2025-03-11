package main

import (
	"github.com/mllukasik/ngw/cmd"
	"os"
)

var version = "latest"

func main() {
	err := cmd.Execute(version)
	if err != nil {
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/alenn-m/interview/svc/cmd"
)

func execute() error {
	return cmd.Execute()
}

func main() {
	if err := execute(); err != nil {
		os.Exit(1)
	}
}

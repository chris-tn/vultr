package main

import (
	"os"

	"github.com/chris-tn/vultr/cmd"
)

func main() {
	cli := cmd.NewCLI()
	cli.RegisterCommands()
	cli.Run(os.Args)
}

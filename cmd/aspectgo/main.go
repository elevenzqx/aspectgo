package main

import (
	"os"

	"github.com/elevenzqx/aspectgo/compiler/cli"
)

func main() {
	os.Exit(cli.Main(os.Args))
}

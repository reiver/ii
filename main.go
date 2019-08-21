package main

import (
	_ "github.com/reiver/ii/commands"
	"github.com/reiver/ii/sys/command"

	"github.com/reiver/go-cli"

	"fmt"
)

func main() {
	fmt.Println("Intergalactic Index")

	var handler cli.Handler = &sys_command.Mux

	cli.RunAndThenExit(handler)
}

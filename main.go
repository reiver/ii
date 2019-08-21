package main

import (
	"github.com/reiver/go-cli"

	"fmt"
)

func main() {
	fmt.Println("Intergalactic Index")

	var handler cli.Handler

	cli.RunAndThenExit(handler)
}

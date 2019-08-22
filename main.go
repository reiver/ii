package main

import (
	_ "github.com/reiver/ii/commands"
	"github.com/reiver/ii/sys/command"

	"github.com/reiver/go-cli"

	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Intergalactic Index (%s) ⚡\n", attn("ii"))

	var handler cli.Handler = &sys_command.Mux

	cli.RunAndThenExit(handler)
}

func attn(s string) string {
	var builder strings.Builder

	builder.WriteString("\x1b[94;40m")
	builder.WriteString(s)
	builder.WriteString("\x1b[0m")

	return builder.String()
}

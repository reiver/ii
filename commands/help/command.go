package verboten

import (
	"github.com/reiver/ii/sys/command"

	"github.com/reiver/go-cli"

	"fmt"
	"io"
	"strings"
)

var (
	name []string = []string{"help"}
)

func init() {
	if err := sys_command.Mux.HandleFunc(run, name...); nil != err {
		panic(err)
	}
}

func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {

	cs, err := sys_command.Mux.Commands()
	if nil != err {
		fmt.Fprintf(stderr, "uh oh: %s\n", err)
		return cli.ExitCodeError
	}

	{
		io.WriteString(stdout, "\nThese are the ii commands:\n")

		var builder strings.Builder

		for _, c := range cs {
			if 0 >= len(c) {
				continue
			}

			builder.Reset()

			builder.WriteString("\n\tii ")
			builder.WriteString(c[0])
			for _, t := range c[1:] {
				builder.WriteRune(' ')
				builder.WriteString(t)
			}
			builder.WriteRune('\n')

			io.WriteString(stdout, builder.String())
		}
	}

	return cli.ExitCodeOK
}

package verboten

import (
	"github.com/reiver/ii/sys/command"
	"github.com/reiver/ii/sys/uhoh"

	"github.com/reiver/go-cli"
	"github.com/reiver/go-iirepo/apps"

	"io"
	"os"
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
		sys_uhoh.Fprintf(stderr, "%s\n", err)
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
			builder.WriteString("\x1b[32m")
			builder.WriteString(c[0])
			builder.WriteString("\x1b[0m")
			for _, t := range c[1:] {
				builder.WriteRune(' ')
				builder.WriteString("\x1b[32m")
				builder.WriteString(t)
				builder.WriteString("\x1b[0m")
			}
			builder.WriteRune('\n')

			io.WriteString(stdout, builder.String())
		}
	}

	var iiapps [][]string
	{
		wd, err := os.Getwd()
		if nil != err {
			sys_uhoh.Fprintf(stderr, "%s\n", err)
			return cli.ExitCodeError
		}

		iiapps, err = iirepo_apps.List(wd)
		if nil != err {
			sys_uhoh.Fprintf(stderr, "%s\n", err)
			return cli.ExitCodeError
		}
	}

	if 0 < len(iiapps) {
		io.WriteString(stdout, "\nThese are the installed ii apps for this repo:\n")

		var builder strings.Builder

		for _, c := range iiapps {
			if 0 >= len(c) {
				continue
			}

			builder.Reset()

			builder.WriteString("\n\tii ")
			builder.WriteString("\x1b[32m")
			builder.WriteString(c[0])
			builder.WriteString("\x1b[0m")
			for _, t := range c[1:] {
				builder.WriteRune(' ')
				builder.WriteString("\x1b[32m")
				builder.WriteString(t)
				builder.WriteString("\x1b[0m")
			}
			builder.WriteRune('\n')

			io.WriteString(stdout, builder.String())
		}
	} else {
		io.WriteString(stdout, "\nThere are no installed ii apps for this repo.\n")
	}

	return cli.ExitCodeOK
}

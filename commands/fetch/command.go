package verboten

import (
	"github.com/reiver/ii/sys/command"
	"github.com/reiver/ii/sys/uhoh"

	"github.com/reiver/go-cli"

	"io"
)

var (
	name []string = []string{"fetch"}
)

func init() {
	if err := sys_command.Mux.HandleFunc(run, name...); nil != err {
		panic(err)
	}
}

func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {

	{
		sys_uhoh.Fprint(stderr, "Command “\x1b[30;41mfetch\x1b[0m” hasn't been implemented yet!\n")
		return cli.ExitCodeError
	}
}

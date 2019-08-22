package verboten

import (
	"github.com/reiver/ii/sys/command"
	"github.com/reiver/ii/sys/uhoh"

	"github.com/reiver/go-cli"
	"github.com/reiver/go-iirepo"

	"fmt"
	"io"
	"os"
)

var (
	name []string = []string{"init"}
)

func init() {
	if err := sys_command.Mux.HandleFunc(run, name...); nil != err {
		panic(err)
	}
}

func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {

	wd, err := os.Getwd()
	if nil != err {
		sys_uhoh.Fprintf(stderr, "%s\n", err)
		return cli.ExitCodeOSError
	}

	if err := iirepo.Init(wd); nil != err {
		sys_uhoh.Fprintf(stderr, "%s\n", err)
		return cli.ExitCodeIOError
	}
	fmt.Fprint(stdout, "initialized\n")

	return cli.ExitCodeOK
}

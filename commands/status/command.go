package verboten

import (
	"github.com/reiver/ii/sys/command"

	"github.com/reiver/go-cli"
	"github.com/reiver/go-iirepo/stage"

	"fmt"
	"io"
	"os"
	"strings"
)

var (
	name []string = []string{"status"}
)

func init() {
	if err := sys_command.Mux.HandleFunc(run, name...); nil != err {
		panic(err)
	}
}

func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {

	io.WriteString(stdout, "Staged files ready to be committed:\n\n")

	wd, err := os.Getwd()
	if nil != err {
		fmt.Fprintf(stderr, "uh oh: %s\n", err)
		return cli.ExitCodeOSError
	}

	fn := func(relstagedpath string) error {
		fmt.Fprintf(stdout, "%s\n", green(relstagedpath))
		return nil
	}

	if err := iirepo_stage.Walk(wd, fn); nil != err {
		fmt.Fprintf(stderr, "uh oh: %s\n", err)
		return cli.ExitCodeIOError
	}

	return cli.ExitCodeOK
}

func green(s string) string {
	var builder strings.Builder
	builder.WriteString("\x1b[32m\t")
	builder.WriteString(s)
	builder.WriteString("\x1b[0m")

	return builder.String()
}

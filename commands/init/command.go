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

	// This is intended to improve the user experience (UX).
	//
	// We check to see if the .ii/ repo exists in the current working directory.
	//
	// And if it already exists, we tell the user about it.
	//
        // Yes, there is a race condition, but this code is more cosmetic from the
        // point of view of correctness.
        //
        // (The race condition is that the .ii/ repo could be deleted between where
        // we check for it and find it here, versus where we try to initialize it.)
	{
		repopath := iirepo.Path(wd)

		fileinfo, err := os.Stat(repopath)
		if nil != err {
			sys_uhoh.Fprintf(stderr, "%s\n", err)
			return cli.ExitCodeOSError
		}

		if fileinfo.IsDir() {
			fmt.Fprint(stdout, "already iniitalized 💚\n")
			return cli.ExitCodeOK
		}
	}

	if err := iirepo.Init(wd); nil != err {
		sys_uhoh.Fprintf(stderr, "%s\n", err)
		return cli.ExitCodeIOError
	}
	fmt.Fprint(stdout, "initialized 💚\n")

	return cli.ExitCodeOK
}

package verboten

import (
	"github.com/reiver/ii/sys/command"

	"github.com/reiver/go-cli"
	"github.com/reiver/go-iirepo"
	"github.com/reiver/go-iirepo/stage"

	"fmt"
	"io"
	"os"
)

var (
	name []string = []string{"drop"}
)

func init() {
	if err := sys_command.Mux.HandleFunc(run, name...); nil != err {
		panic(err)
	}
}

func run(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, command ...string) cli.ExitCode {

	// This is intended to improve the user experience (UX).
	//
	// We check to see if the .ii/ repo exists
	//
	// Yes, there is a race condition, but this code is more cosmetic from the
	// point of view of correctness. (Checks are also done elsewhere.)
	//
	// (The race condition is that the .ii/ repo could be deleted between where
	// we check for it and find it here, versus where we try to drop files to the
	// stage later on.)
	{
		wd, err := os.Getwd()
		if nil != err {
			fmt.Fprintf(stderr, "uh oh: could not get current working directory: %s\n", err)
			return cli.ExitCodeOSError
		}

		_, err = iirepo.LocateRoot(wd)
		if nil != err {
			switch err.(type) {
			case iirepo.NotFound:
				fmt.Fprintf(stderr, "uh oh: no %s/ repository in current directory, or any of the parent directories.\n\nIs that what you expected?\n\nIf you want to create a new ii repo, run:\n\n\tii init\n", iirepo.Name())
			default:
				fmt.Fprintf(stderr, "uh oh: something bad happened when trying to orient myself: %s\n", err)
			}

			return cli.ExitCodeBadConfiguration
		}
	}

	// This too is intended to improve the user experience (UX).
	//
	// We check to see that each file that is suppose to be dropped to the stage, is actualy a file.
	//
	// Yes, there is a race condition, but this is more cosmetic from the point of view of correctness.
	// Checks are also done elsewhere.
	{
		var errored bool

		for _, path := range command {
			isAFile, err := isFile(path)
			if nil != err {
				errored = true
				fmt.Fprintf(stderr, "uh oh: something bad happened when trying to first out if %s is a file: %s\n", path, err)
				continue
			}
			if !isAFile {
				errored = true
				fmt.Fprintf(stderr, "uh oh: cannot drop this, because this is not a file: %s\n", path)
				continue
			}
		}

		if errored {
			return cli.ExitCodeError
		}
	}


	// Try to drop each file to the stage.
	for _, path := range command {
		isAFile, err := isFile(path)
		if nil != err {
			fmt.Fprintf(stderr, "uh oh: something bad happened when trying to first out if %s is a file: %s\n", path, err)
			continue
		}
		if !isAFile {
			fmt.Fprintf(stderr, "uh oh: cannot drop this, because this is not a file: %s\n", path)
			continue
		}

		fmt.Fprintf(stdout, "dropping %s ... ", path)
		if err := iirepo_stage.StoreOriginal(path); nil != err {
			fmt.Fprint(stdout, "error!\n")
			fmt.Fprintf(stderr, "uh oh: something bad happened when trying to stage %s: %s\n", path, err)
			return cli.ExitCodeError
		}
		fmt.Fprint(stdout, "done.\n")
	}

	return cli.ExitCodeOK
}

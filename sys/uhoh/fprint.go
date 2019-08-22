package sys_uhoh

import (
	"fmt"
	"io"
	"strings"
)

const (
	prefix string = "😞 uh oh: "
)

func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, append([]interface{}{prefix}, a...)...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	var builder strings.Builder

	builder.WriteString(prefix)
	builder.WriteString(format)

	return fmt.Fprintf(w, builder.String(), a...)
}

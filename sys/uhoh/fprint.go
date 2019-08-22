package sys_uhoh

import (
	"fmt"
	"io"
	"strings"
)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	var builder strings.Builder

	builder.WriteString("😞 uh oh: ")
	builder.WriteString(format)

	return fmt.Fprintf(w, builder.String(), a...)
}

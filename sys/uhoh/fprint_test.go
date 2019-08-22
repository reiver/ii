package sys_uhoh_test

import (
	"github.com/reiver/ii/sys/uhoh"

	"strings"

	"testing"
)

func TestFprintf(t *testing.T) {

	tests := []struct{
		Format   string
		A      []interface{}
		Expected string
	}{
		{
			Format: "%s",
			A: []interface{}{   "apple"},
			Expected: "😞 uh oh: apple",
		},
		{
			Format: "%s   %s",
			A: []interface{}{   "apple","banana"},
			Expected: "😞 uh oh: apple   banana",
		},
		{
			Format: "%s   %s   %s",
			A: []interface{}{   "apple","banana","cherry"},
			Expected: "😞 uh oh: apple   banana   cherry",
		},
	}

	for testNumber, test := range tests {

		var builder strings.Builder

		sys_uhoh.Fprintf(&builder, test.Format, test.A...)

		if expected, actual := test.Expected, builder.String(); expected != actual {
			t.Errorf("For test #%d, the actual fprint'ed string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

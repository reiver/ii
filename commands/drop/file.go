package verboten

import (
	"os"
)

func isFile(path string) (bool, error) {
	fileinfo, err := os.Stat(path)
	if nil != err {
		return false, err
	}

	if ! fileinfo.Mode().IsRegular() {
		return false, nil
	}

	return true, nil
}

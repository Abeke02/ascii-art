package getData

import (
	"errors"
	"os"
	"strings"
)

func ThirdArgument() (string, error) {
	outFilename := os.Args[3]

	arrOutFile := strings.Split(outFilename, "=")
	if arrOutFile[0] != "--output" || len(arrOutFile) != 2 {
		return "", errors.New("error: incorrect keywords")
	}
	if !strings.HasSuffix(arrOutFile[1], ".txt") {
		return "", errors.New("error: incorrect file name")
	}

	return arrOutFile[1], nil
}

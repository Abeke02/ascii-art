package getData

import (
	"errors"
	"os"
)

func SecondArg() (string, error) {
	style := os.Args[2]
	if style != "standard" && style != "shadow" && style != "thinkertoy" {
		return "", errors.New("error: choose the right style")
	}
	fileStyle := style + ".txt"
	return fileStyle, nil
}

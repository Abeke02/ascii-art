package getData

import (
	"errors"
	"os"
	"strings"
)

func FirstArg() ([]string, error) {
	arg := os.Args[1]
	runes := []rune(arg)
	for i := 0; i < len(runes); i++ {
		if (runes[i] > 127 || runes[i] < 32) && runes[i] != 10 {
			return nil, errors.New("error don't use invalid symbol")
		}
	}

	arg = strings.ReplaceAll(arg, "\\n", "\n")
	arrArg := strings.Split(arg, "\n")
	if arrArg == nil {
		return nil, errors.New("error with argumets")
	}

	count := 0
	for i := 0; i < len(arrArg); i++ {
		if (arrArg[i]) == "" {
			count++
		}
	}

	if count == len(arrArg) {
		slice := arrArg
		arrArg = nil
		for i := 0; i < len(slice); i++ {
			if i != len(slice)-1 {
				arrArg = append(arrArg, slice[i])
			}
		}
	}

	return arrArg, nil
}

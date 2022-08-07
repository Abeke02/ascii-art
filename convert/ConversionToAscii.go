package convert

import (
	"strings"

	"main.go/internal"
)

func ConversionToAscii(words []string, style string) ([]string, error) {
	var slice [][]string
	arr := [1]string{"n"}

	x, err := internal.Width()
	if err != nil {
		return nil, err
	}

	template := internal.ReadtoShare(style)

	count := 0
	for j := 0; j < len(words); j++ {
		for i := 0; i < len(words[j]); i++ {
			slice = append(slice, str(template[words[j][i]-32]))
			count += len(slice[i][0])
			if int(x)-count-5 <= len(str(template[words[j][i]-32])) {
				slice = append(slice, arr[:])
				count = 0
			}
		}
		slice = append(slice, arr[:])
		count = 0
	}

	res := findNewLine(slice)
	return res, nil
}

func str(str string) []string {
	temp := strings.Split(str, "\n")
	return temp
}

func asciiFinal(slice [][]string) string {
	result := ""
	var arr []string
	c := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == nil {
			c += 1
		}
	}

	if c == len(slice) {
		arr = append(arr, "\n")
	} else {
		for j := 0; j < 8; j++ {
			for i := 0; i < len(slice); i++ {
				if slice[i][0] == "" && len(slice[i][6]) == 4 {
					slice[i][6] = "      "
				}
				if slice[i][j] == "" {
					slice[i][j] = "      "
				}

				result += slice[i][j]
			}
			result += "\n"
			arr = append(arr, result)
			result = ""
		}
	}

	for _, i := range arr {
		result += i
	}
	return result
}

func findNewLine(slice [][]string) []string {
	var arr [][]string
	var res []string

	for i := 0; i < len(slice); i++ {
		if slice[i][0] == "n" {
			res = append(res, asciiFinal(arr))
			arr = nil
		} else {
			arr = append(arr, slice[i])
		}
	}
	return res
}

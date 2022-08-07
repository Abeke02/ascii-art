package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateFile(res []string, fileName string) {
	data := ""
	for _, i := range res {
		data += i
	}
	data += "\n"

	if !strings.HasSuffix(fileName, ".txt") {
		fmt.Println("error with file name")
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("error with create file")
		return
	}

	defer file.Close()
	file.Write([]byte(data))
}

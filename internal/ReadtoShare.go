package internal

import (
	"log"
	"os"
	"strings"
)

func ReadtoShare(str string) []string {
	temp, err := os.ReadFile(str)
	if err != nil {
		log.Fatal("file not found \n", err)
		return nil
	}

	if str == "standard.txt" {
		if len(temp) != 6614 {
			log.Fatal("the file is corrupted")
			return nil
		}
	} else if str == "shadow.txt" {
		if len(temp) != 7463 {
			log.Fatal("the file is corrupted")
			return nil
		}
	} else if str == "thinkertoy.txt" {
		if len(temp) != 4703 {
			log.Fatal("the file is corrupted")
			return nil
		}
	}
	template := strings.Split(string(temp), "\n\n")
	return template
}

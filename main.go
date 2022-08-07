package main

import (
	"fmt"
	"os"

	"main.go/convert"
	"main.go/getData"
	"main.go/internal"
)

func main() {
	if len(os.Args[1:]) == 1 {
		haveOne()
	} else if len(os.Args[1:]) == 2 {
		haveTwo()
	} else if len(os.Args[1:]) == 3 {
		haveThree()
	} else {
		fmt.Println("error: incorrect number of arguments")
	}
}

func haveOne() {
	words, err := getData.FirstArg()
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := convert.ConversionToAscii(words, "standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range result {
		fmt.Print(i)
	}
}

func haveTwo() {
	words, err := getData.FirstArg()
	if err != nil {
		fmt.Println(err)
		return
	}
	style, err := getData.SecondArg()
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := convert.ConversionToAscii(words, style)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range result {
		fmt.Print(i)
	}
}

func haveThree() {
	words, err := getData.FirstArg()
	if err != nil {
		fmt.Println(err)
		return
	}
	style, err := getData.SecondArg()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName, err := getData.ThirdArgument()
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := convert.ConversionToAscii(words, style)
	if err != nil {
		fmt.Println(err)
		return
	}
	internal.CreateFile(result, fileName)
}

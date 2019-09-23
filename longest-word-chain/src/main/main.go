package main

import (
	"fmt"
	"longest-word-chain/src/file"
)

func main() {
	fmt.Println("hello world!")
	strArr := []string {"123","fuck","go"}
	file.WriteFileAndCreate("test",strArr)
	lines,err := file.RedaDate("test")
	if err != nil {
		fmt.Println("error")
		return
	}
	for _, item := range lines {
		fmt.Println(item)
	}
}

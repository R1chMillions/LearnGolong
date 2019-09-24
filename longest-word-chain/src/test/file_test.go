package test

import (
	"fmt"
	"longest-word-chain/src/file"
	"testing"
)

func TestFile(t *testing.T) {
	fmt.Println("hello world!")
	strArr := []string{"123", "fuck", "go"}
	file.WriteFileAndCreate("test", strArr)
	lines, err := file.RedaDate("test")
	if err != nil {
		fmt.Println("error")
		return
	}
	for item := lines.Front(); nil != item; item = item.Next() {
		fmt.Printf("%s", item.Value)
	}
}

func TestReadData(t *testing.T) {
	lines, _ := file.RedaDate("/workspaces/golong/longest-word-chain/data")
	for item := lines.Front(); item != nil; item = item.Next() {
		fmt.Printf("%s", item.Value)
	}
}

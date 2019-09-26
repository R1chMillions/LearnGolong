package test

import (
	"container/list"
	"fmt"
	"longest-word-chain/src/file"
	"testing"
)

func TestFile(t *testing.T) {
	list := list.New()
	list.PushBack("123")
	list.PushBack("test")
	list.PushBack("go")
	file.WriteFileAndCreate("test", list)
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

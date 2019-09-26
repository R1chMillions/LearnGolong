package file

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteFileAndCreate(fileName string, list *list.List) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Create File. Name: %s\n", fileName)
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Printf("Create failed error: %s\n", err)
			return
		}
		fmt.Printf("File '%s' Created!\n", fileName)
	}

	fmt.Println("Writing data...")
	for item := list.Front(); item != nil; item = item.Next() {
		file.WriteString(item.Value.(string))
		file.WriteString(" ")
	}
	fmt.Println("Done")
}

func RedaDate(fileName string) (*list.List, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("read failed")
		return nil, err
	}
	var l = list.New()
	bfRd := bufio.NewReader(file)
	for {
		line, err := bfRd.ReadBytes('\n')
		str := string(line)
		if str == "" || str == "\n" {
			continue
		}
		l.PushBack(strings.ToLower(strings.Replace(str, "\n", "", -1)))
		if err == io.EOF {
			return l, nil
		} else if err != nil {
			return l, err
		}
	}
}

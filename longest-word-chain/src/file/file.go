package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFileAndCreate(fileName string, dataArr []string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Create File. Name:%s", fileName)
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println("Create failed error:%s", err)
			return
		}
		fmt.Println("File '%s' Created!", fileName)
	}

	fmt.Println("Writing data...")
	for _, item := range dataArr {
		file.WriteString(item)
	}
	fmt.Println("Done")
}


func RedaDate(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("read failed")
		return nil, err
	}
	var strArr []string
	bfRd := bufio.NewReader(file)
	var index = 0
	strArr = make([]string,5)
	for {
		line, err := bfRd.ReadBytes('\n')
		strArr[index] = string(line)
		index++
		if err != nil {
			if err == io.EOF {
				return strArr, nil
			}
			return strArr, err
		}
	}
	return strArr, err
}

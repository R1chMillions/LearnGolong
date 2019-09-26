package main

import "longest-word-chain/src/file"

func main() {
	path := "/workspaces/golong/longest-word-chain/"
	lines, _ := file.RedaDate(path + "data")
	res := Do(lines)
	file.WriteFileAndCreate(path+"answer", res)
}

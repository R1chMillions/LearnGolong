package main

import "longest-word-chain/src/file"

func main() {
	lines, _ := file.RedaDate("/workspaces/golong/longest-word-chain/data")
	Do(lines)
}

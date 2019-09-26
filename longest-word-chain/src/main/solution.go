package main

import (
	"container/list"
	"fmt"
	"longest-word-chain/src/file"
	"longest-word-chain/src/plist"
)

func Do(l *list.List) {
	hc, ec := GetWordHeaderAndEnd(l)
	graph := MakeGraph(hc, ec, l.Len())

	chans := list.New()
	for i := 0; i < len(graph); i++ {
		chanList := list.New()
		chanList.PushBack(plist.GetListValue(l, i))
		visitedArr := make([]bool, l.Len())
		visitedArr[i] = true
		for j := 0; j < len(graph); j++ {
			if graph[i][j] == 1 && i != j && visitedArr[j] != true {
				visitedArr[j] = true
				chanList.InsertBefore(plist.GetListValue(l, j), chanList.Back())
			}
		}
		for k := 0; k < len(graph); k++ {
			if graph[k][i] == 1 && i != k && visitedArr[k] != true {
				visitedArr[k] = true
				chanList.InsertAfter(plist.GetListValue(l, k), chanList.Front())
			}
		}
		if chanList.Front() != nil {
			chans.PushBack(chanList)
		}
	}
	PrintMax(l, chans)
	PrintAll(l, chans)
}

func MakeGraph(headerChars []byte, endChars []byte, lenth int) [][]int {
	graph := make([][]int, lenth)
	for i := 0; i < lenth; i++ {
		graph[i] = make([]int, lenth)
		for j := 0; j < lenth; j++ {
			graph[i][j] = 0
		}
	}
	for i := 0; i < len(headerChars); i++ {
		for j := 0; j < len(endChars); j++ {
			if headerChars[i] == endChars[j] && i != j {
				graph[i][j] = 1
			}
		}
	}

	return graph
}

func GetWordHeaderAndEnd(l *list.List) ([]byte, []byte) {
	headerArr := make([]byte, l.Len())
	endArr := make([]byte, l.Len())
	i := 0
	for item := l.Front(); item != nil; item = item.Next() {
		s := item.Value.(string)
		headerArr[i] = s[0]
		endArr[i] = s[len(s)-1]
		i++
	}
	return headerArr, endArr
}

func PrintAll(l *list.List, chans *list.List) {
	fmt.Print("\n\nPrint all:\n")
	i := 0
	for chan1 := chans.Front(); chan1 != nil; chan1 = chan1.Next() {
		fmt.Printf("%d.", i)
		for item := chan1.Value.(*list.List).Front(); item != nil; item = item.Next() {
			fmt.Printf("%s ", item.Value)
		}
		fmt.Println()
		i++
	}
}

func PrintMax(l *list.List, chans *list.List) {
	lenth := 0
	var maxChan *list.List
	for chan1 := chans.Front(); chan1 != nil; chan1 = chan1.Next() {
		curr := chan1.Value.(*list.List)
		if lenth < curr.Len() {
			lenth = curr.Len()
			maxChan = curr
		}
	}

	file.WriteFileAndCreate("answer", maxChan)

	fmt.Print("\n\nPrint max:\n")
	for item := maxChan.Front(); item != nil; item = item.Next() {
		fmt.Printf("%s ", item.Value)
	}
}

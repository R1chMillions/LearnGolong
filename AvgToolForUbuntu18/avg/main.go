package main

import(
	"fmt"
	"os"
	"strconv"
)

func GetAvg(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	avg := sum/len(s)
	return avg
}

func main()  {
	var s []int 
	osArgs := os.Args[1:]
	for i := 0; i < len(osArgs); i++ {
		item, err := strconv.Atoi(osArgs[i])
		if err != nil {
			fmt.Println("strconv err!")
			fmt.Println(err)
			return
		}
		s = append(s, item)
	}
	fmt.Println(GetAvg(s))
}
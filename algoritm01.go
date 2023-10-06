package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	a := os.Args
	arr := make([]int, 0)

	var i int64 = 0

	for ; i < 20; i++ {
		convert := strconv.FormatInt(i, 10)
		fmt.Println(i)

		if convert == a[len(a)-1] {
			fmt.Println("if is runned")
			continue
		}
		arr = append(arr, int(i))
		fmt.Println(arr)
	}

}
func push_notU() {

}

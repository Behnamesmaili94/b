package main

import (
	"fmt"
)

type a struct {
	x []int
}

func main() {
	fmt.Println("app is running")
	var m a
	m.lappend(0, 10, 9)
	m.x[1] = m.x[len(m.x)-1]
	m.x = m.x[:len(m.x)-1]
	fmt.Println(m.x)

}

func (a *a) lappend(start int, end int, value int) {
	for i := start; i < end; i++ {
		a.x = append(a.x, value)
		value++
	}

}

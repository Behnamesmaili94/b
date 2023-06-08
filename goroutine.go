package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main is running : \n")

	//first coreone run after go coreone run then go coretwo run
	//when complate coreone core two running without go coretwo
	
	go coreOne()
	go coreTwo()

	coreOne()

	coreTwo()

}

func coreOne() {

	for a := 0; a < 10; a++ {

		runtime.Gosched()
		fmt.Println("core 1 is running ", a)
	}

}

func coreTwo() {

	for a := 0; a < 10; a++ {

		runtime.Gosched()
		fmt.Println("core 2 is running ", a)
	}
}

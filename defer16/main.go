package main

import "fmt"

func main() {
	defer fmt.Println("two")
	defer fmt.Println("one")
	defer fmt.Println("world")
	fmt.Println("Hello")
	myDefer()

}

// world, One, two
//0, 1, 2, 3, 4
// hello, 43210, two, One, world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
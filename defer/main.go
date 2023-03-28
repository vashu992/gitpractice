package main

import "fmt"

func main() {

	for i := 1; i < 11; i++ {
	  defer fmt.Println(i)
	}
// defer koi gadbadi ho to sambhal lo
	defer fmt.Println("hello")
	defer fmt.Println("world")

	fmt.Println("koi bhi message")
	defer fmt.Println("ashutosh")
	
	fmt.Println("koi bhi message-----")

}
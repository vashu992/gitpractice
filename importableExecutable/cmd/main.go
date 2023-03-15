package main

import (
	
	"fmt"

	"github.com/vashu992/gitpractice/importableExecutable/calculator"
)

func main() {

	usr := calculator.User{}
	usr.Name ="ghgjkhghh"
	//fmt.Println("sum = ", calculator.add(54646,5465))
	fmt.Println("Multiply = ", calculator.Multiply(54646,5465))
	fmt.Println("Divide = ", calculator.Divide(54646,5465))
	fmt.Println("sub = ", calculator.Sub(54646,5465))
}
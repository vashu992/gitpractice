package main

import (
	"fmt"

	"github.com/vashu992/gitpractice/importableExecutable/calculator"
	"github.com/vashu992/gitpractice/importableExecutable/store"
)

func main() {

	usr := calculator.User{}
	usr.Name = "ghgjkhghh"
	//fmt.Println("sum = ", calculator.add(54646,5465))
	fmt.Println("Multiply = ", calculator.Multiply(54646, 5465))
	fmt.Println("Divide = ", calculator.Divide(54646, 5465))
	fmt.Println("sub = ", calculator.Sub(54646, 5465))

	// import using interface
	var dboprs store.DBOperations
	dboprs = store.Store{}
	dboprs.SaveRecord(" just print this as record ....")
}

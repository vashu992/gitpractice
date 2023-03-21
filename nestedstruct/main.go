package main

import"fmt"


func main() {
	fmt.Println("Student struct")
    // Student struct in golang
	
	ashutosh := Student{"ashutosh",2,3,4,12}
	fmt.Println(ashutosh)
	fmt.Println("Student details are %+v\n Student")

    vishal := Teacher{"vishal",5,"golang",34}
	fmt.Println(vishal)
	fmt.Println("Teacher details are %+v\n Techer")
}

type Student struct{
Name string
Roll int
Class int
Room int
Age int
}

// created nested struct

type Teacher struct{
	Name string
	exp int 
	subject string
	Age int
}


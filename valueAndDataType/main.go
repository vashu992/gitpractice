package main

import "fmt"

func main() {
	// 12 -> int
	// 12.22 -> float
	// true/false -> bool
	// c -> char / string
	// git practice -> String
    const number2 = 45645

	var number int
	number = -12
    number= 5645464
	var addrsOfNumber *int
    addrsOfNumber = &number

	number3:=5454
	fmt.Println(number3)


	//num4 := number
	
	var decNum float32
	decNum =13211.3554151
    var addOfFloat * float32
    addOfFloat = &decNum

    var flags bool
	flags = true
    var addOfFlag *bool
	addOfFlag = &flags

	var name string
	name = "gitpractice"
	var addOfName  *string
	addOfName = &name

	fmt.Printf("Number value = % , decimal value = %v ,  flags value = %v ,name value = %s \n ",number , decNum , flags , name)
	fmt.Printf("Address ofNumber value = %v , add of decimal value = %v ,  address of flags value = %v ,name value = %s \n",addrsOfNumber , addOfFloat , addOfFlag , *addOfName)
    

}
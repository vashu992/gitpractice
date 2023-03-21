package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	ashutosh := User{"Ahutosh","ashutoshpratap992@gmail.com",true, 32}
	fmt.Println(ashutosh)
	fmt.Println("ashutosh details are %+v\n", ashutosh)
	fmt.Println("Name is %v and emails is %v.",ashutosh.Name,ashutosh.Email)
}

type User struct {
	Name  string
	Email  string
	Status  bool
	Age     int
}
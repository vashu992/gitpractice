package main

import "fmt"

type UserA struct {
	Name      string
	Address   string
 	ContactNo int
}

type UserB struct {
	Name      string
	Address   string
	ContactNo int
}

type UserOperations interface{
	addUser()
}

func main() {

	var vishal UserA
	vishal.Name = "Vishal Singh"
	vishal.Address = "Varanasi"
	vishal.ContactNo = 12324355432
	 shivam := UserB {
		Name: "Shivam Singh",
		Address: "same hai Varanasi",
		ContactNo: 5678548754897,
	}

	// vishal.addUser()
	// shivam.addUser()
	
    var UserOperations UserOperations
	UserOperations = vishal // User A
	UserOperations.addUser()

	UserOperations = shivam
	UserOperations.addUser()
}


func (user UserA) addUser()  {
	user.Name = "763426734673476437843"
	fmt.Println("User a wala method chal raha hai , ",user)
}



func (user UserB) addUser() {
	fmt.Println("User B wala method chal raha hai , ",user)
}
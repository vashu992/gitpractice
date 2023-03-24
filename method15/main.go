package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	ashutosh := User{"Ahutosh","ashutoshpratap992@gmail.com",true, 32}
	fmt.Println(ashutosh)
	fmt.Println("ashutosh details are %+v\n", ashutosh)
	fmt.Println("Name is %v and emails is %v.\n",ashutosh.Name,ashutosh.Email)
	ashutosh.GetStatus()
	ashutosh.NewMail()
}

type User struct {
	Name  string
	Email  string
	Status  bool
	Age     int
}

func ( u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "ashutoshpratap1998@gmail.com"
	fmt.Println("Email of this user is ", u.Email)
}
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
	vishal.addUser(5)
	fmt.Println(" after add User vishal = ", vishal)

	vishal.addUserDub(54646766)
	fmt.Println("value from addDub = ", vishal)



	shivam.addUser(5)
	add(66787)

}

func add(a int ) {
	fmt.Println("function run ho raha ho raha hai, aur a = " , a)
}
func (user UserA) addUser(a... int)  {
	user.Name = "763426734673476437843"
	fmt.Println("User a wala method chal raha hai , ",user)
}

func (u *UserA) addUserDub(a int) int {
	u.Name = "Name value change kar raha hu"
	fmt.Println("User a dub wala method chal raha hai , ",u)
	return 0
}
func (user UserB) addUser(a... int) {
	fmt.Println("User B wala method chal raha hai , ",user)
}
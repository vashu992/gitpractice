package main

import "fmt"


// non primitive data type

type Student struct{

	Name string
	Class int
	RollNumber string
    StudentAddress Address

}

type Address struct{
   Lane1 string
   Lane2 string
   Post string
   Dist string
   Village string
}
type Phone struct {
    BasicInfo
	IMEI string
    Config
}
type Laptop struct{
    BasicInfo
	SerialNo string
	Configuration Config
}
type BasicInfo struct{
    Brand string
	Model string
}
type Config struct{
    RAM int
	ROM int
	Processor string
	OS string 
}	



func main() {
	//Primitive approch
    // var name vishal string
   //  var class vishal int
   //  var rollnumber_vishal int

	// var name shivam string
	// var class shivam int
	// var rollnumber_shivam int

//Non- Primitive
  var vishal Student
  vishal.Class = 12
  vishal.Name = "Vishal Singh"
  vishal.StudentAddress.Dist = "Bhadohi"
  vishal.StudentAddress.Village ="Bhaisahata"

   shivam := Student{
	Name: "Shivam Singh",
	Class: 12,
	RollNumber: "50",
    StudentAddress: Address{
	Dist: "Distict val",	
	Village: "Village val",
	Lane1: "NH2",
	},
   }

   val := 122
   val2 :="12212"

   var interfaceExample interface{}
   interfaceExample = val

   fmt.Println("Interface value = ",interfaceExample)

   interfaceExample = val2
   fmt.Println("Interface value = ",interfaceExample)

   interfaceExample = false
   fmt.Println("Interface value = ",interfaceExample)


	fmt.Println("hello+++++++ Team...this is Vishal's",vishal)
	fmt.Println("hello+++++++ Team...this is Shivam's",shivam)
}
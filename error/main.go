package main

import (
	"fmt"
	"os"
)

func main() {

	// // system generated error
	// b := []int{
	// 	1, 2, 3, 4,
	// }
	// fmt.Println(b)
	// fmt.Println(len(b))
	// fmt.Println(b[4])


	// user generated error
    
// fatel error (it stops execution)
dbconnection := true
if !dbconnection{
	fatelerror("Can not connect with db so closing Programe")

}
pass := "1234"

fmt.Println(len(pass))
   if len(pass) < 8 {
	   err := fmt.Errorf("password dose not match criteria ,password = %v",pass)
	 
	  panic(err)
	}else{
	   savetodb(pass)
	}
}

func fatelerror(message string) {
	fmt.Printf("some error occure which is causing fatel error , message = %v\n",message)
	os.Exit(1)
}


func savetodb(data string) {
	fmt.Println("Record saved success, record =",data)
}
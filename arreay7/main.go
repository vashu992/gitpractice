package main

import "fmt"

func main() {
	fmt.Println("welcome to the conding concepts in array class")

	var  fruitList [8]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"
	fruitList[3] = "Orange"
	fruitList[4] = "Papaya"
	fruitList[5] = "Carrot"
	fruitList[6] = "Monkey"
	fruitList[7] = "Animal"

	fmt.Println("Fruit list: ", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

		
	var vegList = [5]string{"potato","beans","mushroom"}
	fmt.Println("Vegy list is: ",len(vegList))
}
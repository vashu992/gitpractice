package main

import "fmt"

func main() {
	fmt.Println("conding concepts by golang in loops")

	days := []string{"Sunday","Monday", "Tuesday", "Wendesday", "Friday","Saturaday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days{
	// 	fmt.Println("index is %v and value is %v\n", day)
	// }


	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			break
		}
		fmt.Println("value is: ", rougueValue)
		rougueValue++

	}

lco:
     fmt. Println("Jumoing conding concepts663")

}
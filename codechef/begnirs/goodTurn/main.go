package main

import "fmt"

func main() {
	var t int
	//testcase
	fmt.Scanf("%d\n",&t)
	for t > 0 {
		x , y := 0,0
		fmt.Scanf("%d",&x)
		fmt.Scanf("%d",&y)
		if (x+y) > 6 {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}

		t--
	}
}
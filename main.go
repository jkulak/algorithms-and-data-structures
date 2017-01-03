package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x)
	var y *int

	fmt.Println(y)

	var z = new(int)
	fmt.Println(z)
	*z = 10
	fmt.Println(z)
	fmt.Println(*z)
	fmt.Println(&z)

}

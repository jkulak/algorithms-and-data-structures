package main

import "fmt"

func main() {

	// Arrays - have defined size
	// * var a1 []int - wrong declaration, memory space not assigned
	// * unset values are set to 0s

	// One dimensional array of int
	var a1 [3]int
	a1[0] = 1
	a1[2] = 1
	fmt.Printf("%v\n", a1)

	a11 := [3]int{91, 92, 93}
	fmt.Printf("%v\n", a11)

	// One dimensional array with initialization
	var a2 = [3]int{1, 2, 3}
	fmt.Printf("%v\n", a2)

	// Initialization with make keywrod
	var a3 []int
	a3 = make([]int, 3)
	a3[1] = 9
	fmt.Printf("%v\n", a3)

	// Double dimensional array of int
	var a4 [3][3]int
	a4[0][0] = 1
	a4[1][1] = 2
	a4[2][2] = 3
	fmt.Printf("%v\n", a4)

	// Double dimensional array of int with initialization
	var a5 = [4][3]int{[3]int{9}, [3]int{0, 6, 0}, [3]int{0, 0, 3}}
	fmt.Printf("%v\n", a5)

	var a6 [3][3]int
	a6[0] = [3]int{4}
	a6[1] = [3]int{0, 6}
	a6[2] = [3]int{0, 0, 8}
	fmt.Printf("%v\n", a6)

	a7 := [3][3]int{[...]int{1, 2, 3}, [...]int{1, 2, 3}, [...]int{1, 2, 3}}
	fmt.Printf("%v\n", a7)

	a8 := [1][2]int{[...]int{5, 2}}
	fmt.Printf("%v\n", a8)

	// Arrays have their place, but they're a bit inflexible, so you don't see
	// them too often in Go code. Slices, though, are everywhere. They build on
	// arrays to provide great power and convenience.
	// The type specification for a slice is []T, where T is the type of the
	// elements of the slice. Unlike an array type, a slice type has no
	// specified length.

	// Copy
	// b = make([]T, len(a))
	// copy(b, a) // or b = append([]T(nil), a...)
	//
	// Cut
	// a = append(a[:i], a[j:]...)
	//
	// Delete
	// a = append(a[:i], a[i+1:]...) // or a = a[:i+copy(a[i:], a[i+1:])]
	//
	// Delete without preserving order
	// a[i], a = a[len(a)-1], a[:len(a)-1]
	//
	// Pop
	// x, a = a[len(a)-1], a[:len(a)-1]
	//
	// Push
	// a = append(a, x)

	s1 := []string{"a", "b", "c"}
	fmt.Printf("%v\n", s1)

	// Creating a slice using make
	// func make([]T, len, cap) []T
	s2 := make([]string, 10, 11)
	s2[4] = "HELLO"
	fmt.Printf("%v\n", s2)

	// Create a slice from an array
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s3 := x[:] // a slice referencing the storage of x
	fmt.Printf("%v\n", s3)

	// Append elements to a slice
	s4 := []byte{1}
	s4 = append(s4, 99, 200, 201, 255)
	fmt.Printf("%v\n", s4)

	// Append slice to a slice
	s5 := []byte{4, 5, 6, 7}
	s6 := []byte{44, 45, 46, 47}
	s5 = append(s5, s6...)
	s5 = append(s5, []byte{1, 1, 1, 1}[:2]...)
	fmt.Printf("%v\n", s5)

	// Two dimensioanl slice (slice of slices)
	// Allocate the top-level slice.
	sa, sb := 4, 4
	s7 := make([][]uint8, sa) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range s7 {
		s7[i] = make([]uint8, sb)
	}
	fmt.Printf("%v\n", s7)
}

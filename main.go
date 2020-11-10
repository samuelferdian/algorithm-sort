package main

import (
	"fmt"
)

func swap(x, y int) (int, int) {
	return y, x
}

func printResult(s string, r []int) {
	fmt.Printf("%s %v\n", s, r)
}

func main() {
	//================//
	// Selection sort //
	//================//
	fmt.Println("SELECTION SORT")
	s := []int{64, 34, 25, 12, 22, 11, 90}
	printResult("Before:", s)

	for i := 0; i < len(s); i++ {
		low := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[low] {
				low = j
			}
		}
		if s[i] > s[low] {
			s[i], s[low] = swap(s[i], s[low])
		}
	}
	printResult("After:", s)

	//=============//
	// Bubble sort //
	//=============//
	fmt.Println("BUBBLE SORT")
	b := []int{64, 34, 25, 12, 22, 11, 90}
	printResult("Before:", b)

	for i := 0; i < len(s); i++ {
		swapped := false
		for j := 0; j < len(b)-i-1; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = swap(b[j], b[j+1])
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	printResult("After:", b)

	//================//
	// Insertion sort //
	//================//
	fmt.Println("INSERTION SORT")
	in := []int{64, 34, 25, 12, 22, 11, 90}
	printResult("Before:", in)
	for i := 1; i < len(in); i++ {
		curr := in[i]
		back := i - 1

		for (back >= 0) && (in[back] > curr) {
			in[back+1] = in[back]
			back = back - 1
		}
		in[back+1] = curr
	}
	printResult("After:", in)

}

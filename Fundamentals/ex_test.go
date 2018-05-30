package Fundamentals

import (
	"fmt"
)

func ExampleCompare3Ints() {
	fmt.Println(Compare3Ints(1, 2, 3))
	fmt.Println(Compare3Ints(3, 3, 3))
	// Output:
	// false
	// true
}

func ExampleBetween01() {
	fmt.Println(Between01(0.4, 0.7))
	fmt.Println(Between01(1, 0.7))
	fmt.Println(Between01(0, 3))
	// Output:
	// true
	// false
	// false
}

func ExampleInt2BinStr() {
	fmt.Println(Int2BinStr(4))
	// Output:
	// 100
}

func ExamplePrintBools() {
	bools := [][]bool{
		{true, true, false, false},
		{false, true, false, true, true, true},
		{true, false, true, false},
		{true, true, true, false},
	}
	PrintBools(bools)
	// Output:
	// h 0 1 2 3 4 5
	// 0 * * X X
	// 1 X * X * * *
	// 2 * X * X
	// 3 * * * X
}

func ExamplePrintTMatrix() {
	m := [][]int{
		{2, 3, 4, 5},
		{9, 8, 7, 6},
	}
	PrintTMatrix(m)
	// Output:
	//    2   9
	//    3   8
	//    4   7
	//    5   6
}

func ExampleLg() {
	fmt.Println(Lg(1024))
	fmt.Println(Lg(1025))
	fmt.Println(Lg(1023))
	// Output:
	// 10
	// 10
	// 9
}

func ExampleHistogram() {
	a := []int{1, 1, 2, 3, 1, 7, 5, 3, 2, 2, 2}
	fmt.Println(Histogram(a, 8))
	fmt.Println(Sum(Histogram(a, 8)) == len(a))
	// Output:
	// [0 3 4 2 0 1 0 1]
	// true
}

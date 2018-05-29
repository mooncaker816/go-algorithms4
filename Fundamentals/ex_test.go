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

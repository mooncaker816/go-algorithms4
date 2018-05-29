package Fundamentals

// Compare3Ints 比较三个整数是否相等
// Ex 1.1.3
func Compare3Ints(a, b, c int) bool {
	// a, err := fmt.Scanf("%d\n", os.Stdin)
	// if err != nil {
	// 	log.Fatalf("can not read from stdin: %v\n", err)
	// }
	// b, _ := fmt.Scanf("%d\n", os.Stdin)
	// if err != nil {
	// 	log.Fatalf("can not read from stdin: %v\n", err)
	// }
	// c, _ := fmt.Scanf("%d\n", os.Stdin)
	// if err != nil {
	// 	log.Fatalf("can not read from stdin: %v\n", err)
	// }
	if a == b {
		if b == c {
			return true
		}
	}
	return false
}

// Between01 检查 x,y 是否属于（0，1）
// Ex 1.1.5
func Between01(x, y float64) bool {
	if x >= 1 || x <= 0 {
		return false
	}
	if y >= 1 || y <= 0 {
		return false
	}
	return true
}

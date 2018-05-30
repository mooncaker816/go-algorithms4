package Fundamentals

import (
	"fmt"
	"strconv"
	"strings"
)

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

// Int2BinStr 将整数转为二进制字符串
// Ex 1.1.9
func Int2BinStr(n int) string {
	return fmt.Sprintf("%b\n", n)
}

// PrintBools 打印二维布尔数组
func PrintBools(bs [][]bool) {
	var b strings.Builder
	b.WriteString("h ")
	maxcolumn := 0
	for _, row := range bs {
		if len(row) > maxcolumn {
			maxcolumn = len(row)
		}
	}
	for i := 0; i < maxcolumn; i++ {
		b.WriteString(strconv.Itoa(i))
		if i < maxcolumn-1 {
			b.WriteString(" ")
		}
	}
	b.WriteString("\n")
	for i, bi := range bs {
		b.WriteString(strconv.Itoa(i))
		b.WriteString(" ")
		for j, bj := range bi {
			if bj {
				b.WriteString("*")
			} else {
				b.WriteString("X")
			}
			if j < len(bi)-1 {
				b.WriteString(" ")
			}
		}
		b.WriteString("\n")
	}
	fmt.Println(b.String())
}

// PrintTMatrix 打印 a 的转置
func PrintTMatrix(a [][]int) {
	for j := 0; j < len(a[0]); j++ {
		for i := 0; i < len(a); i++ {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Println()
	}
}

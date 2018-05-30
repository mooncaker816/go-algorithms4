package Fundamentals

import (
	"fmt"
	"math"
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
// Ex 1.1.13
func PrintTMatrix(a [][]int) {
	for j := 0; j < len(a[0]); j++ {
		for i := 0; i < len(a); i++ {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Println()
	}
}

// Lg 计算不大于logn的最大整数
// Ex 1.1.14
func Lg(n int) int {
	p := uint(0)
	for {
		if 1<<p > n {
			break
		}
		p++
	}
	return int(p - 1)
}

// Histogram Ex 1.1.15
func Histogram(a []int, m int) []int {
	mp := make(map[int]int)
	ret := make([]int, m)
	for _, v := range a {
		mp[v]++
	}
	for i := 0; i < m; i++ {
		if v, ok := mp[i]; ok {
			ret[i] = v
		} else {
			ret[i] = 0
		}
	}
	return ret
}

// Fibo 计算第 n 个斐波那契项
// Ex 1.1.19
func Fibo(n int) int {
	f1, f2 := 1, 1
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

// LnN 计算 ln(n!)
// Ex 1.1.20
func LnN(n uint) float64 {
	if n <= 1 {
		return 0
	}
	return math.Log(float64(n)) + LnN(n-1)
}

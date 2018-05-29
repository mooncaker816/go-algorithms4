package Fundamentals

// GCD returns greatest common divisor of p and q
// 欧几里得算法求解非负整数p,q的最大公约数
func GCD(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return GCD(q, r)
}

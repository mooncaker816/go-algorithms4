package Fundamentals

// Average 计算a中元素的平均值
func Average(a []float64) float64 {
	sum := 0.0
	for _, f := range a {
		sum += f
	}
	return sum / float64(len(a))
}

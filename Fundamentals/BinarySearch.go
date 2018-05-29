package Fundamentals

// Rank 二分查找法寻找 a 中值为 key 的秩,a 已升序排列
func Rank(key int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// RankR 递归版二分查找法
func RankR(key int, a []int) int {
	return rankr(key, a, 0, len(a)-1)
}

func rankr(key int, a []int, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	if key < a[mid] {
		return rankr(key, a, lo, mid-1)
	}
	if key > a[mid] {
		return rankr(key, a, mid+1, hi)
	}
	return mid
}

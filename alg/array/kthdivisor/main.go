package kthdivisor

func kthFactor(n int, k int) int {
	l := 0
	r := n
	i := 0
	for l < r {
		l++
		if n%l != 0 {
			continue
		}
		i++
		if i == k {
			return l
		}
		r = n / l
	}
	return -1
}

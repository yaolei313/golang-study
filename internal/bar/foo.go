package bar

func Fibonacci() func() int {
	count := 0
	pre0 := 0
	pre1 := 1
	sum := 0
	return func() int {
		if count == 0 {
			sum = 0
		} else if count == 1 {
			sum = 1
		} else {
			sum = pre0 + pre1
			pre0, pre1 = pre1, sum
		}
		count += 1
		return sum
	}
}

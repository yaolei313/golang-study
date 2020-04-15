package bar

func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

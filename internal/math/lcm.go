package math

// https://en.wikipedia.org/wiki/Least_common_multiple

func Lcm(first int, integers []int) int {
	result := first * integers[0] / Gcd(first, integers[0])
	for i := 1; i < len(integers); i++ {
		result = Lcm(result, []int{integers[i]})
	}

	return result
}

// https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	//
	return a
}

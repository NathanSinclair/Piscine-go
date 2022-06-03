package piscine

// Takes an int min and an int max as parameters and returns a slice of ints with all the values between the min and max, without using the append function
func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	ans := make([]int, max-min)
	count := 0
	for i := min; i < max; i++ {
		ans[i-min] = min + count
		count++
	}
	return ans
}

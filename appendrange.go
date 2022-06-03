package piscine

// Takes an int min and an int max as parameters and returns the values between the min and max, without using the make function
func AppendRange(min, max int) []int {
	var ans []int
	for i := min; i < max; i++ {
		ans = append(ans, i)
	}
	return ans
}

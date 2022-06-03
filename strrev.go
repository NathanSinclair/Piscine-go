package piscine

// Reverses a string and returns the reversed string
func StrRev(s string) string {
	i := 0
	l := len(s) - 1
	str := []byte(s)
	retstr := []byte(s)
	for range s {
		retstr[i] = str[l]
		i++
		l--
	}
	return string(retstr)
}

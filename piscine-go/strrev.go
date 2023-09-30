package piscine

func StrRev(s string) string {
	str := []rune(s)
	for i := 0; i <= (len(str)-1)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	return string(str)
}

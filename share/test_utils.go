package share

//GenerateStringForLength this function return one string with length equal of (length int) param
func GenerateStringForLength(length int) string {
	c := []rune{}
	for i := 0; i < length; i++ {
		t := []rune("a")
		c = append(c, t[0])
	}

	return string(c)
}

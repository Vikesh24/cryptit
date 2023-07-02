package decrypt

func Nimbus(str string) string {
	decryptStr := ""

	for _, i := range str {
		ascii := int(i)
		character := string(rune(ascii - 3))
		decryptStr += character
	}

	return decryptStr
}
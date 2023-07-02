package encrypt

func Nimbus(uStr string) string {
	encryptStr := ""

	for _, i := range uStr {
		ascii := int(i)
		character := string(rune(ascii + 3))
		encryptStr += character
	}

	return encryptStr
}
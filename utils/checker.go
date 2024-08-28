package check

func CheckIn(str string) bool {
	for _, char := range str {
		if (char < 32 || char > 126)  && (char != '\r'  && char != '\n'){
			// fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
			return false
		}
	}
	return true
}

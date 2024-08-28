package fs

func Printer(inputLine string, slice [][]string) string {
	var s string
	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			index := int(char) - 32
			s += slice[index][j]
		}
		s += "\r\n"
	}
	return s
}

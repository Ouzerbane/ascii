package fs

// gets rid of the first empty string if the input has ONLY newlines
func OnlyNewLine(inputLines []string) []string {
	onlyNewLine := true
	for _, value := range inputLines {
		if value != "" {
			onlyNewLine = false
			break
		}
	}
	if onlyNewLine {
		inputLines = inputLines[1:]
	}
	return inputLines
}

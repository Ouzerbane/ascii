package fs

import "strings"

var Slice [][]string

func Splitter(banner, str string) {
	var lines []string
	if banner == "thinkertoy" {
		str = strings.ReplaceAll(str, "\r\n", "\n")
	}
	lines = strings.Split(str, "\n\n")
	for i := range lines {
		if lines[i] != "" {
			// gets rid of the first new line
			if i == 0 {
				lines[0] = lines[0][1:]
			}
			Slice = append(Slice, strings.Split(lines[i], "\n"))
		}
	}
}

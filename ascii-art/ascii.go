package ascii

import (
	"strings"

	fs "youmed/ascii-art/funcs"
)

func Ascii(banner string, input string) string {
	var s string
	str := fs.Reader(banner)
	fs.Splitter(banner, str)

	inputLines := strings.Split(input, "\r\n")
	// checking if the input has only newlines
	inputLines = fs.OnlyNewLine(inputLines)
	for _, value := range inputLines {
		if value == "" {
			s += "\r\n"
		} else {
			s += fs.Printer(value, fs.Slice)
		}
	}
	fs.Slice = [][]string{}
	return s
}

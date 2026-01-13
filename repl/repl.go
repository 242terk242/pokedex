package repl

import (
	"strings"
)

func CleanInput(text string) []string {
	var r []string
	trim_string := strings.TrimSpace(text)
	text_string := strings.Fields(trim_string)

	for _, str := range text_string {
		r = append(r, strings.ToLower(str))
	}

	return r
}

package repl

import (
	"strings"
)

func cleanInput(text string) []string {
	var r []string
	text_string := strings.Split(text, `,`)

	for _, str := range text_string {
		if str != "" {
			r = append(r, strings.ToLower(str))
		}
	}

	return r
}

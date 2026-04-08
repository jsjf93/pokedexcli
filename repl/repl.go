package repl

import "strings"

func cleanInut(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

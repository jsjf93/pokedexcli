package repl

import "strings"

func CleanInut(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

package main

import (
	"strings"
)

func format(s string, length int) string {
	if len(s) > length {
		return s[:length]
	}
	return s + strings.Repeat(" ", length-len(s))
}
package main

import "strings"

func joinStrings(elements ...string) string {
	return strings.Join(elements, "-")
}

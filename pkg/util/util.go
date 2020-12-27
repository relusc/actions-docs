package util

import (
	"strings"
)

// Contains checks if a given slice contains a given entry
func Contains(s []string, name string) bool {
	for _, item := range s {
		if item == name {
			return true
		}
	}
	return false
}

// Normalize removes all spaces from s to get readable and nicely indented content
func Normalize(s string) string {
	var normalized string
	segments := strings.Split(s, "\n")
	for _, segment := range segments {
		normalized += strings.TrimSpace(segment) + "\n"
	}
	return normalized
}

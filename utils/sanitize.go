package utils

import "regexp"

func Sanitize(str string, strip string) string {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z]+`)
	nonAlphaRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	if strip == "alpha" {
		return nonAlphaRegex.ReplaceAllString(str, "")
	} else if strip == "alphanumeric" {
		return nonAlphanumericRegex.ReplaceAllString(str, "")
	}
	return ""
}

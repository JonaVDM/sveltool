package utils

import "strings"

func FirstLetterUpper(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func FirstLetterLower(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

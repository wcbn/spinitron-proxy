package api

import (
	"regexp"
)

// TODO harden these regexes to handle query strings and other edge cases

func IsResourcePath(path string) bool {
	re, _ := regexp.Compile(`\/api\/\w+\/\d+.*`)
	return re.MatchString(path)
}

func IsCollectionPath(path string) bool {
	re, _ := regexp.Compile(`\/api\/\w+.*`)
	return re.MatchString(path)
}

package api

import (
	"net/url"
	"regexp"
	"strings"
)

// TODO harden these, test for edge cases

func IsResourcePath(path string) bool {
	re, _ := regexp.Compile(`\/api\/\w+\/\d+.*`)
	return re.MatchString(path)
}

func IsCollectionPath(path string) bool {
	if IsResourcePath(path) {
		return false
	}
	re, _ := regexp.Compile(`\/api\/\w+.*`)
	return re.MatchString(path)
}

func GetCollectionName(s string) string {
	dummy := "https://foo.com/"
	u, _ := url.Parse(dummy + s)
	segments := strings.Split(u.Path, "/")

	for i := range segments {
		if segments[i] == "api" || segments[i] == "" {
			continue
		}
		return segments[i]
	}

	return ""
}

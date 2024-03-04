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
	re, _ := regexp.Compile(`\/api\/\w+.*`)
	return re.MatchString(path)
}

func GetCollectionName(pathWithQuery string) string {
	dummy := "https://foo.com/"
	u, _ := url.Parse(dummy + pathWithQuery)
	segments := strings.Split(u.Path, "/")

	for i := range segments {
		if segments[i] == "api" || segments[i] == "" {
			continue
		}
		return segments[i]
	}

	return ""
}

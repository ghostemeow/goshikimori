package utils

import (
	"net/url"
	"unicode"
)

// DEPRECATED (issue #42)
//
// url.QueryEscape() breaks Cyrillic with spaces, and without it,
// Latin with spaces breaks, so there is this intermediate function
// that solves this problem and there are no more errors.
func LanguageCheck(target string) string {
	for i := 0; i < len(target); i++ {
		if target[i] > unicode.MaxASCII {
			return target
		}
	}
	return url.QueryEscape(target)
}

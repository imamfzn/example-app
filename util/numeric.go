package util

import "regexp"

func Numeric(s string) bool {
	return regexp.MustCompile(`^\d*$`).MatchString(s)
}

package entity

import "strings"

var providerPrefixes = map[string][]string{
	"telkomsel": {"+62852", "+62811", "+62821"},
	"xl":        {"+62878", "+62817"},
	"tri":       {"+62896"},
}

// number should in E164 format
func Provider(number string) string {
	for provider, prefixes := range providerPrefixes {
		for _, prefix := range prefixes {
			if strings.HasPrefix(number, prefix) {
				return provider
			}
		}
	}
	return ""
}

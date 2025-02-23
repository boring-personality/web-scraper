package utils

import "net/url"

func ResolveURL(base string, href string) string {
	parsedBase, _ := url.Parse(base)
	parsedHref, err := url.Parse(href)
	if err != nil {
		return ""
	}
	return parsedBase.ResolveReference(parsedHref).String()
}

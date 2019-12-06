package string

import "net/url"

func IsValidUrl(urls string) bool {
	_, err := url.ParseRequestURI(urls)
	if err != nil {
		return false
	} else {
		return true
	}
}

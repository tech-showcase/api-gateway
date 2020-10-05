package helper

import (
	"net/url"
)

func JoinURL(baseUrl *url.URL, pathUrlStr string) (joinedUrl *url.URL, err error) {
	pathUrl, err := url.Parse(pathUrlStr)
	if err != nil {
		return
	}

	joinedUrl = baseUrl.ResolveReference(pathUrl)
	return
}

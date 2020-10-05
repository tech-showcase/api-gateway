package helper

import (
	"net/http"
	"net/url"
	"time"
)

func JoinURL(baseUrl *url.URL, pathUrlStr string) (joinedUrl *url.URL, err error) {
	pathUrl, err := url.Parse(pathUrlStr)
	if err != nil {
		return
	}

	joinedUrl = baseUrl.ResolveReference(pathUrl)
	return
}

func GetQueryStringValue(r *http.Request, key string) (value string) {
	if valueArr, ok := r.URL.Query()[key]; ok {
		value = valueArr[0]
	}

	return
}

func ParseDateTime(timeStr string) (value time.Time, err error) {
	value, err = time.Parse(time.RFC3339Nano, timeStr)
	if err != nil {
		return
	}

	return
}

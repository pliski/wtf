package urlcheck

import (
	"net/url"
)

const InvalidResultCode = 999

type uriResult struct {
	url           string
	resultCode    int
	resultMessage string
	valid         bool
}

func newUriResult(urlString string) *uriResult {

	uResult := uriResult{
		url: urlString,
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		if len(urlString) == 0 {
			uResult.resultMessage = "empty url"
			uResult.resultCode = InvalidResultCode
			uResult.valid = false
		} else {
			uResult.resultMessage = "invalid url"
			uResult.resultCode = InvalidResultCode
			uResult.valid = false
		}
	} else {
		uResult.valid = true
	}

	return &uResult
}

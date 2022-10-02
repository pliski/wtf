package urlcheck

import (
	"net/url"
)

const InvalidResultCode = 999

type urlResult struct {
	url           string
	resultCode    int
	resultMessage string
	isValid       bool
}

func newUrlResult(urlString string) *urlResult {

	uResult := urlResult{
		url: urlString,
	}

	if len(urlString) == 0 {
		uResult.resultMessage = "Empty url"
		uResult.resultCode = InvalidResultCode
		uResult.isValid = false
		return &uResult
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		uResult.resultMessage = "Invalid url"
		uResult.resultCode = InvalidResultCode
		uResult.isValid = false
		return &uResult
	}

	uResult.isValid = true
	return &uResult
}

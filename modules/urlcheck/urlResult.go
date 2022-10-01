package urlcheck

import (
	"net/url"
)

type uriResult struct {
	raw           *url.URL
	setting       string
	resultCode    StatusMsg
	resultMessage string
	valid         bool
	// httpClient    *http.Client
	get fUC
}

func newUriResult(urlString string) *uriResult {

	uResult := uriResult{
		setting: urlString,
	}

	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		if len(urlString) == 0 {
			uResult.resultMessage = "empty url"
			uResult.resultCode = 0
			uResult.valid = false
		} else {
			uResult.resultMessage = "invalid url"
			uResult.resultCode = 0
			uResult.valid = false
		}
	} else {
		// uResult.httpClient = &http.Client{Timeout: 20 * time.Second}
		uResult.get = CheckSomeUrl()
		uResult.raw = u // TODO: not needed?
		uResult.valid = true
	}

	return &uResult
}

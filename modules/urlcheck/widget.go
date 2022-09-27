package ipinfo

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type uriResult struct {
	raw           *url.URL
	setting       string
	resultCode    int
	resultMessage string
	valid         bool
}

type Widget struct {
	view.TextWidget

	// result   string
	settings *Settings
	uriList  []uriResult
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	maxUri := len(settings.paramList)

	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
		uriList:  make([]uriResult, maxUri),
	}

	widget.View.SetWrap(false)
	widget.sanitize()

	return &widget
}

func (widget *Widget) Refresh() {
	widget.check()

	// widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, widget.result, false })

	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {

	content := ""
	for _, ur := range widget.uriList {

		content += fmt.Sprintf("%s: [%d] %s", ur.setting, ur.resultCode, ur.resultMessage)
	}

	return widget.CommonSettings().Title, content, true
}

// this method reads the config and calls ipinfo for ip information
func (widget *Widget) check() {

	for _, ur := range widget.uriList {
		if ur.valid {
			c := &http.Client{Timeout: 10 * time.Second} // TODO: parametric timeout
			res, err := c.Get(ur.setting)
			if err != nil {
				ur.resultMessage = err.Error()
			} else {
				ur.resultCode = res.StatusCode
				ur.resultMessage = http.StatusText(res.StatusCode)
			}
		}
	}
}

func (widget *Widget) sanitize() {

	for _, line := range widget.settings.paramList {

		uResult := uriResult{
			setting: line,
		}

		u, err := url.ParseRequestURI(line)
		if err != nil {
			if len(line) == 0 {
				// fmt.Println("Received an empty line")
				uResult.resultMessage = "empty"
				uResult.resultCode = 0
				uResult.valid = false
			} else {
				// fmt.Println("Invalid uri:", line)
				uResult.resultMessage = "invalid"
				uResult.resultCode = 0
				uResult.valid = false
			}
		} else {
			uResult.raw = u
			uResult.valid = true
		}
		// Store the uri
		widget.uriList = append(widget.uriList, uResult)
	}
}

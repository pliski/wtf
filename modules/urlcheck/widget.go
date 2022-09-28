package urlcheck

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

	settings *Settings
	uriList  []uriResult
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	maxUri := len(settings.paramList)

	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings: settings,
		uriList:  make([]uriResult, maxUri),
	}

	widget.View.SetWrap(false)
	widget.sanitize()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.check()
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}

func (widget *Widget) content() string {

	content := ""
	for _, ur := range widget.uriList {
		content += fmt.Sprintf("%s: [%d] %s\n", ur.setting, ur.resultCode, ur.resultMessage)
	}

	return content
}

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
				uResult.resultMessage = "empty"
				uResult.resultCode = 0
				uResult.valid = false
			} else {
				uResult.resultMessage = "invalid"
				uResult.resultCode = 0
				uResult.valid = false
			}
		} else {
			uResult.raw = u
			uResult.valid = true
		}

		widget.uriList = append(widget.uriList, uResult)
	}
}

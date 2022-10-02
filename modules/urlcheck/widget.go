package urlcheck

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
	uriList  []*uriResult
	client   *http.Client
	timeout  time.Duration
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	maxUri := len(settings.urls)

	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings: settings,
		uriList:  make([]*uriResult, maxUri),
		client:   GetClient(),
		timeout:  time.Duration(settings.requestTimeout) + time.Second,
	}

	widget.View.SetWrap(false)
	widget.init()

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
		if ur.resultCode == 0 && ur.valid {
			ur.resultMessage = "wait..."
		}
		// content += fmt.Sprintf("%s: [%d] %s\n", ur.setting, ur.resultCode, ur.resultMessage)
		content += fmt.Sprintf("%s: %s\n", ur.url, ur.resultMessage)
	}
	return content
}

func (widget *Widget) check() {
	for _, urlRes := range widget.uriList {
		if urlRes.valid {
			urlRes.resultCode, urlRes.resultMessage = DoRequest(urlRes.url, widget.timeout, widget.client)
		}
	}
}

func (widget *Widget) init() {
	for i, urlString := range widget.settings.urls {
		widget.uriList[i] = newUriResult(urlString)
	}
}

package urlcheck

import (
	"fmt"
	"net/http"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
	uriList  []*uriResult
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	logger.Log("[urlcheck] New Widget")
	maxUri := len(settings.paramList)

	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings: settings,
		uriList:  make([]*uriResult, maxUri),
	}

	widget.View.SetWrap(false)
	widget.init()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	logger.Log("[urlcheck] Refresh")
	widget.check()
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	logger.Log("[urlcheck] Display")
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}

func (widget *Widget) content() string {
	logger.Log("[urlcheck] Content")

	content := ""
	for _, ur := range widget.uriList {
		logger.Log(fmt.Sprintf("[urlcheck] Content %s: %d", ur.setting, ur.resultCode))
		if ur.resultCode == 0 && ur.valid {
			ur.resultMessage = "wait..."
		}
		content += fmt.Sprintf("%s: [%d] %s\n", ur.setting, ur.resultCode, ur.resultMessage)
	}

	return content
}

func (widget *Widget) check() {
	logger.Log("[urlcheck] Check")
	logger.Log(fmt.Sprintf("[urlcheck] Check() widget address: %p", widget))
	logger.Log(fmt.Sprintf("[urlcheck] Check() urilist address: %p", widget.uriList))
	logger.Log(fmt.Sprintf("[urlcheck] Check() urilist[1] address: %p", &widget.uriList[1]))

	for i, urlRes := range widget.uriList {
		if urlRes.valid {
			// newUrlRes := urlRes                          // TODO urilist must be an array of pointes to urlResult
			// c := &http.Client{Timeout: 10 * time.Second} // TODO: parametric timeout
			// res, err := c.Get(urlRes.setting)
			// res, err := urlRes.httpClient.Get(urlRes.setting)
			statusCode, err := urlRes.get(*urlRes)
			if err.Critical {
				logger.Log(fmt.Sprintf("[urlcheck] GET %s failed: %s", urlRes.setting, err.Error()))
				urlRes.resultMessage = err.Error()
			} else {
				logger.Log(fmt.Sprintf("[urlcheck] GET %s: %d", urlRes.setting, statusCode))
				urlRes.resultCode = statusCode
				urlRes.resultMessage = http.StatusText(int(statusCode))
			}
			widget.uriList[i] = urlRes
			logger.Log(fmt.Sprintf("[urlcheck] GET %s: %d", urlRes.setting, statusCode))
		}
	}
}

func (widget *Widget) init() {
	for i, urlString := range widget.settings.paramList {
		widget.uriList[i] = newUriResult(urlString)
	}
}

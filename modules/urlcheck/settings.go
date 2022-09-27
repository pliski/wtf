package ipinfo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "urlcheck"
)

type Settings struct {
	*cfg.Common

	paramList []string `help:"A list of uri to check"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	settings.paramList = cfg.ParseAsMapOrList(ymlConfig, "paramList")

	settings.SetDocumentationPath("urlcheck")

	return &settings
}

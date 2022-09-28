package urlcheck

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "urlcheck"
)

type Settings struct {
	common *cfg.Common

	paramList []string `help:"A list of uri to check"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}
	settings.paramList = cfg.ParseAsMapOrList(ymlConfig, "paramList")

	return &settings
}

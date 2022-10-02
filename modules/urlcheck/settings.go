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

	requestTimeout int      `help:"Max Request duration in seconds"`
	urls           []string `help:"A list of url to check"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		requestTimeout: ymlConfig.UInt("timeout", 30),
	}
	settings.urls = cfg.ParseAsMapOrList(ymlConfig, "urls")
	return &settings
}

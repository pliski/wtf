package urlcheck

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "urlcheck"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	paramList []string `help:"A list of uri to check"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}
	settings.paramList = cfg.ParseAsMapOrList(ymlConfig, "paramList")

	// settings.SetDocumentationPath("urlcheck")

	return &settings
}

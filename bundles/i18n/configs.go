package i18n

import (
	"gopkg.in/orivil/orivil.v0"
)

var Config = &struct {
	Languages           map[string]string
	DefaultLang         string
	CookieKey           string
	Auto_generate_files bool
}{
	// set default value
	Languages: map[string]string{
		"简体中文":    "zh-CN",
		"english": "en",
	},
	DefaultLang:         "english",
	CookieKey:           "i18n-language",
	Auto_generate_files: true,
}

// sortName: fullName
var langs map[string]string

// use shortName for get fullName
func GetFullName(shortName string) string {
	return langs[shortName]
}

func init() {

	orivil.Cfg.ReadStruct("i18n.yml", Config)
	langs = make(map[string]string, len(Config.Languages))
	for fullName, shortName := range Config.Languages {
		langs[shortName] = fullName
	}
}

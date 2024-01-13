package config

import (
	"WePanel/backend/global"
	"github.com/Unknwon/goconfig"
)

func GetConfig(section string, key string) string {
	cfg, err := goconfig.LoadConfigFile(global.CONFIGDIR)
	if err != nil {
		global.LOG.Fatalf("Unable to load configuration file：%s", err)
	}
	result, err := cfg.GetValue(section, key)
	if err != nil {
		global.LOG.Errorf("Load [%s] config failed：%s", section, err)
	}
	return result
}

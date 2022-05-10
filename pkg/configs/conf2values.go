package configs

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var (
	AppSetting = &App{}
	cfg        *ini.File
	err        error
)

// Conf2Values 读取配置变量
func Conf2Values() {
	cfg, err = ini.Load("configs/app.ini")
	if err != nil {
		log.Fatalf("pkg/configs/con2values.Conf2Values: %v", err)
	}

	mapTo("app", AppSetting)
}

// mapTo 转换配置变量
func mapTo(section string, v interface{}) {
	err = cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("pkg/configs/con2values.mapTo %s err: %v", section, err)
	}
}

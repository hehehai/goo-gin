package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

//项目配置信息初始化
var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	//加载配置文件
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// 获取默认配置的运行模式属性，设置默认值
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//server 配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(9090)
	ReadTimeout = sec.Key("READ_TIMEOUT").MustDuration(60) * time.Second
	WriteTimeout = sec.Key("WRITE_TIMEOUT").MustDuration(60) * time.Second
}

//应用配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

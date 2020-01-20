package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	MAIL_DRIVER       string
	MAIL_HOST         string
	MAIL_PORT         int
	MAIL_USERNAME     string
	MAIL_PASSWORD     string
	MAIL_ENCRYPTION   string
	MAIL_FROM_ADDRESS string
	MAIL_FROM_NAME    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadMail()
}


func LoadMail() {
	sec, err := Cfg.GetSection("mail")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	MAIL_DRIVER = sec.Key("MAIL_DRIVER").MustString("stmp")
	MAIL_HOST = sec.Key("MAIL_HOST").MustString("127.0.0.1")
	MAIL_PORT = sec.Key("MAIL_PORT").MustInt(465)
	MAIL_USERNAME = sec.Key("MAIL_USERNAME").MustString("")
	MAIL_PASSWORD = sec.Key("MAIL_PASSWORD").MustString("")
	MAIL_ENCRYPTION = sec.Key("MAIL_ENCRYPTION").MustString("")
	MAIL_FROM_ADDRESS = sec.Key("MAIL_FROM_ADDRESS").MustString("")
	MAIL_FROM_NAME = sec.Key("MAIL_FROM_NAME").MustString("")
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!#)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

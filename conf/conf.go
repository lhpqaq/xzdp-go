package conf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env string

	Hertz Hertz `yaml:"hertz"`
	MySQL MySQL `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
	Cors  Cors  `yaml:"cors"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type Hertz struct {
	Address         string `yaml:"address"`
	EnablePprof     bool   `yaml:"enable_pprof"`
	EnableGzip      bool   `yaml:"enable_gzip"`
	EnableAccessLog bool   `yaml:"enable_access_log"`
	LogLevel        string `yaml:"log_level"`
	LogFileName     string `yaml:"log_file_name"`
	LogMaxSize      int    `yaml:"log_max_size"`
	LogMaxBackups   int    `yaml:"log_max_backups"`
	LogMaxAge       int    `yaml:"log_max_age"`
}
type CORSWhitelist struct {
	AllowOrigin      string `yaml:"allow-origin"`
	AllowMethods     string `yaml:"allow-methods"`
	AllowHeaders     string `yaml:"allow-headers"`
	ExposeHeaders    string `yaml:"expose-headers"`
	AllowCredentials bool   `yaml:"allow-credentials"`
}

type Cors struct {
	Mode      string          `yaml:"mode"`
	WhiteList []CORSWhitelist `yaml:"white_list"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	_, confpath, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("error"))
	}
	prefix := confpath[0 : len(confpath)-7]
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	fmt.Println(confFileRelPath)
	content, err := os.ReadFile(confFileRelPath) // replace the deprecated ioutil.ReadFile to os.ReadFile
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()

	//pretty.Printf("%+v\n", conf)
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}

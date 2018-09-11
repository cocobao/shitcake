package conf

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/cocobao/log"

	yaml "gopkg.in/yaml.v2"
)

var GCfg *Config

type stLog struct {
	LogDir   string `yaml:"log_dir"`
	LogLevel int    `yaml:"log_level"`
}

type Config struct {
	Port       string `yaml:"port"`
	MongoDb    string `yaml:"mongo_db"`
	Log        stLog  `yaml:"log"`
	StaticPath string `yaml:"static_path"`
	ViewPath   string `yaml:"view_path"`
}

func ParseConfig(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read config file fail")
		os.Exit(0)
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		fmt.Println("unmarshal config fail, path:", path)
		os.Exit(0)
	}
	return cfg
}

func setupLogging(path string) {
	log.NewLogger(GCfg.Log.LogDir, GCfg.Log.LogLevel)
	log.Debug("setup log ok")
}

func SetupConfig() {
	GCfg = ParseConfig("conf/setting.yaml")
	setupLogging("conf/log.xml")
}

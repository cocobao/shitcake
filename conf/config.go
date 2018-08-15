package conf

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/cocobao/log"

	yaml "gopkg.in/yaml.v2"
)

var GCfg *Config

type Config struct {
	Port       string `yaml:"port"`
	MongoDb    string `yaml:"mongo_db"`
	StaticPath string `yaml:"static_path"`
	ViewPath   string `yaml:"view_path"`
}

func Unmarshal(path string) *Config {
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

	log.Debug("setup log ok")
}

func SetupConfig() {

	GCfg = Unmarshal("conf/setting.yaml")
}

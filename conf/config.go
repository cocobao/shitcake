package conf

import (
	"fmt"
	"os"

	"io/ioutil"

	log "github.com/cihub/seelog"

	yaml "gopkg.in/yaml.v2"
)

var GCfg *Config

type Config struct {
	Port    string `yaml:"port"`
	MongoDb string `yaml:"mongo_db"`
}

func Unmarshal(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("read config file fail")
		os.Exit(0)
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		fmt.Println("unmarshal config fail, path:", path)
		os.Exit(0)
	}
	return cfg
}

func SetupConfig() {
	GCfg = Unmarshal("conf/setting.yaml")
}

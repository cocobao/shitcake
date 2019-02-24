package server

import (
	"github.com/cocobao/shitcake/conf"
	"github.com/cocobao/shitcake/routers"
	"github.com/cocobao/shitcake/store"
)

func Setup() {
	conf.SetupConfig()
	store.SetupMongoDB(conf.GCfg.MongoDb)
	routers.Run(conf.GCfg.Port)
}

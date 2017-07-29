package server

import (
	"github.com/cocobao/shitcake/conf"
	"github.com/cocobao/shitcake/routers"
)

func Setup() {
	conf.SetupConfig()
	// store.SetupMongoDB(conf.GCfg.MongoDb)
	routers.Run()
}

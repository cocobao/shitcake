package main

import (
	"github.com/astaxie/beego"
	"github.com/cocobao/shitcake/conf"
	_ "github.com/cocobao/shitcake/routers"
	"github.com/cocobao/shitcake/store"
)

func main() {
	conf.SetupConfig()
	store.SetupMongoDB(conf.GCfg.MongoDb)
	beego.Run()
}

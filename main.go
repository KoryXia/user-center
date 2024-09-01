package main

import (
	log "github.com/sirupsen/logrus"
	"user-center/global"
	"user-center/initialize"
)

func main() {
	initialize.Log()
	log.Info("Starting user-center...")
	global.DB = initialize.DB()
	global.Cache = initialize.Cache()
	initialize.ServerAndRouters()
}

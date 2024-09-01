package main

import (
	"fmt"
	"user-center/global"
	"user-center/initialize"
)

func main() {
	fmt.Println("MySQL initialized")
	global.DB = initialize.DB()
	initialize.ServerAndRouters()
}

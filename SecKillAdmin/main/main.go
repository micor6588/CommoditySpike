package main

import (
	_ "CommoditySpike/SecKillAdmin/router"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}

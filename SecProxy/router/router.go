package router

import (
	"CommoditySpike/SecProxy/controller"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.Debug("enter router init")
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill") //实现抢购接口的路由
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo") //当前秒杀状态的路由
}

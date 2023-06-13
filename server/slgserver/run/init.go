package run

import (
	"slgserver/config"
	"slgserver/db"
	"slgserver/net"
	"slgserver/server/slgserver/controller"
	"slgserver/server/slgserver/logic"
	"slgserver/server/slgserver/logic/mgr"
	"slgserver/server/slgserver/model"
	"slgserver/server/slgserver/static_conf"
	"slgserver/server/slgserver/static_conf/facility"
	"slgserver/server/slgserver/static_conf/general"
	"slgserver/server/slgserver/static_conf/npc"
	"slgserver/server/slgserver/static_conf/skill"
)

var MyRouter = &net.Router{}

func Init() {
	db.TestDB()

	static_conf.Basic.Load()
	static_conf.MapBuildConf.Load()
	static_conf.MapBCConf.Load()

	facility.FConf.Load()
	general.GenBasic.Load()
	skill.Skill.Load()
	general.General.Load()
	npc.Cfg.Load()

	serverId := config.File.MustInt("logic", "server_id", 1)
	model.ServerId = serverId

	logic.BeforeInit()

	mgr.NMMgr.Load()
	//需要先加载联盟相关的信息
	mgr.UnionMgr.Load()
	mgr.RAttrMgr.Load()
	mgr.RCMgr.Load()
	mgr.RBMgr.Load()
	mgr.RFMgr.Load()
	mgr.RResMgr.Load()
	mgr.SkillMgr.Load()
	mgr.GMgr.Load()
	mgr.AMgr.Load()
	logic.Init()
	logic.AfterInit()

	//全部初始化完才注册路由，防止服务器还没启动就绪收到请求
	initRouter()
}

func initRouter() {

	controller.DefaultRole.InitRouter(MyRouter)
	controller.DefaultMap.InitRouter(MyRouter)
	controller.DefaultCity.InitRouter(MyRouter)
	controller.DefaultGeneral.InitRouter(MyRouter)
	controller.DefaultArmy.InitRouter(MyRouter)
	controller.DefaultWar.InitRouter(MyRouter)
	controller.DefaultCoalition.InitRouter(MyRouter)
	controller.DefaultInterior.InitRouter(MyRouter)
	controller.DefaultSkill.InitRouter(MyRouter)
}

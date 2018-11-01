package router

import (
	"sync"

	"github.com/sifbc/pledge/apiserver/handler"

	"github.com/gin-gonic/gin"
)

// Router 全局路由
var router *gin.Engine
var onceCreateRouter sync.Once

func GetRouter() *gin.Engine {
	onceCreateRouter.Do(func() {
		router = createRouter()
	})

	return router
}

func createRouter() *gin.Engine {
	router := gin.Default()
	// 版本控制
	// v1 := Router.Group("/v1")
	// {
	// factor
	bill := router.Group("/pledge")
	{
		bill.POST("/uploadgoodsinfo", handler.UploadPledgeInfo)        //上传押品信息
		bill.POST("/monitorReception", handler.UploadPatrolInfo)       //上传巡库信息
		bill.POST("/insuranceReception", handler.UploadInsurNotarInfo) //上传保险和公证信息
		bill.POST("/uploadWarningInfo", handler.UploadWarningInfo)
		bill.POST("/statusSync", handler.StatusSync)

		bill.POST("/periodSet", handler.SetAlertPeriod) //设置预警周期

		bill.POST("/querygoodsinfo", handler.QueryGoodsInfo)
		bill.POST("/queryperiod", handler.QueryPeriod)
		bill.POST("/querymonitorpledge", handler.QueryMonitorPledge)

		//Probe Information
		bill.HEAD("/keepaliveQuery", handler.KeepaliveQuery)
	}

	//测试
	test := router.Group("/scfs/blockChain")
	{
		test.POST("/pledgeReception", handler.PledgeReception) //押品信息接收接口
	}
	//upload schema json file
	return router
}

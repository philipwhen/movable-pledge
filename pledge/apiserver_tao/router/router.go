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
		bill.POST("/uploadgoodsinfo", handler.UploadGoodsInfo)
		bill.POST("/uploadpatrolinfo", handler.UploadPatrolInfo)
		bill.POST("/uploadWarningInfo", handler.UploadWarningInfo)
		bill.POST("/statusSync", handler.StatusSync)
		bill.POST("/querygoodsinfo", handler.QueryGoodsInfo)

		//Probe Information
		bill.HEAD("/keepaliveQuery", handler.KeepaliveQuery)
	}
	//upload schema json file
	return router
}

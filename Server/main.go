/*
Version: 0.01
Author: Mang Zhang, Shenzhen China
Release Date: 2025-10-19
Project Name: GoK8sCalluna
Description: A tool to help mange K8s.
Copy Rights: MIT License
Email: m13692277450@outlook.com
Mobile: +86-13692277450
HomePage: www.pavogroup.top

*/

package main

import (
	loginControllers "gok8s/controllers/login"
	"gok8s/kubernetsServ"
	"gok8s/models"
	"gok8s/routers"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

//	func middleWear(c *gin.Context) {
//		c.Next()  // 处理请求, 继续执行主程序的后续的中间件或处理函数，然后再回来
//		c.Abort() // 中断请求，不再执行后续的中间件或处理函数,直接向下执行
//		log.Println("中间件执行完成")
//	}

func main() {
	go loginControllers.InitDB()
	//go config.InitDB()
	kubernetsServ.K8sConnectionInit()
	kubernetsServ.GetK8sPods()
	kubernetsServ.GetK8sResources()
	// 初始化Kubernetes相关数据
	kubernetsServ.GetPVCList()
	//kubernetsServ.Deployment()
	//go kubernetsServ.Deployment()

	r := gin.Default()
	// 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	//Need put before LoadHTMLGlob
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	r.LoadHTMLGlob("templates/**/*")

	r.Static("/static", "./static")

	r.StaticFS("/website", http.Dir("./website"))

	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.LoginRoutersInit(r)
	routers.K8sManageCenterInit(r)
	routers.K8sResourcesInit(r)
	routers.PodsLogRouterInit(r)
	//routers.PVCControllerInit(r)
	// 初始化控制器
	loginController := loginControllers.LoginController{}
	kubernetsServ.GetK8sResources()
	// 路由配置 - 明确区分GET和POST
	r.GET("/login", loginController.ShowLoginPage)
	r.POST("/login", loginController.Login)

	//r.Use(middleWear, middleWear) // 全局中间件, 多个中间件用逗号隔开
	//如果使用了goroutine，则必须使用c.context拷贝,c.Copy()
	r.Run(":8080")

}

// UnixToTime 将 Unix 时间戳转换为格式化的时间字符串

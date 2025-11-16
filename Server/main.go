/*
Version: 0.01
Author: Mang Zhang, Shenzhen China
Release Date: 2025-10-19
Project Name: GoK8sCalluna
Description: A tool to help manage K8s.
Copy Rights: MIT License
Email: m13692277450@outlook.com
Mobile: +86-13692277450
HomePage: www.pavogroup.top , github.com/13692277450

*/

package main

import (
	"embed"
	"gok8s/config"
	"gok8s/kubernetsServ"
	"gok8s/models"
	"gok8s/routers"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

//	func middleWear(c *gin.Context) {
//		c.Next()  // 处理请求, 继续执行主程序的后续的中间件或处理函数，然后再回来
//		c.Abort() // 中断请求，不再执行后续的中间件或处理函数,直接向下执行
//		log.Println("中间件执行完成")
//	}
//
// go: embed: ./templates/**/*
var templateFiles embed.FS

func main() {
	// 读取模板文件
	// data, err := fs.ReadFile(templateFiles, "templates/example.html")
	// if err != nil {
	// 	// 处理错误
	// }
	config.LogrusConfigInit()
	time.Sleep(400 * time.Millisecond)
	config.Lg.Info("GoK8s backend server started... ")

	//go loginControllers.InitDB()
	//go config.InitDB()
	kubernetsServ.K8sConnectionInit()
	if kubernetsServ.Clientset == nil {
		config.Lg.Info("Kubernetes client initialization failed")
	}

	// Get PVCs
	kubernetsServ.GetPVCList(kubernetsServ.Namespace)

	// Get K8s resources
	kubernetsServ.GetK8sResources()

	// Get K8s Pods
	kubernetsServ.GetK8sPods()
	r := gin.Default()
	// 增强的CORS中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24小时
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

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
	// r.LoadHTMLGlob("templates/**/*")
	// r.Static("/static", "./static")
	// r.StaticFS("/website", http.Dir("./website"))
	// Initialize routers
	routers.DeployYamlRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.LoginRoutersInit(r)
	routers.K8sManageCenterRoutersInit(r)
	routers.K8sResourcesRoutersInit(r)
	routers.PodsLogRoutersInit(r)
	routers.SystemCenterRoutersInit(r)
	routers.NamespaceRoutersInit(r)
	routers.PVCRoutersInit(r)
	routers.MetricsRoutersInit(r)
	routers.ServicesRoutersInit(r)

	// Check routers registration
	// routes := r.Routes()
	// for _, route := range routes {
	// 	fmt.Printf("Registered route: %s %s\n", route.Method, route.Path)
	// 	if strings.Contains(route.Path, "deploypod") {
	// 		fmt.Printf("FOUND DEPLOYPOD ROUTE: %s %s\n", route.Method, route.Path)
	// 	}
	// }

	//
	// loginController := loginControllers.LoginController{}
	//utils.GetNodeLogs("k8s-master01")
	//utils.CaptureNodeExecOutput1(false, "./kubernetsServ/kubeconfig", "k8s-master01", []string{"ls"}) //, " get", " pods"})
	//utils.GetNodeLogs("k8s-master01")

	// GET and POST
	// r.GET("/login", loginController.ShowLoginPage)
	// r.POST("/login", loginController.Login)
	//kubernetsServ.GetPodResources("default", "mysqlwordpress")
	// go services.K8sNodesPerformance()
	// go services.K8sPodsPerformance()
	//r.Use(middleWear, middleWear) // 全局中间件, 多个中间件用逗号隔开
	//如果使用了goroutine，则必须使用c.context拷贝,c.Copy()
	r.Run(":8080")

}

// UnixToTime 将 Unix 时间戳转换为格式化的时间字符串

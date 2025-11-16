package routers

import (
	"gok8s/controllers/deploy" // 直接导入deploy包

	"github.com/gin-gonic/gin"
)

func DeployYamlRoutersInit(r *gin.Engine) {
	// 使用/api作为基础路径
	yamlRouters := r.Group("/api")
	{
		// 注册DeployYamlController来处理YAML部署
		yamlRouters.POST("/deploypod", deploy.DeployYamlController)
		yamlRouters.POST("/deploynamespace", deploy.DeployNamespaceControllers{}.DeployNamespaceController)
		// fmt.Println("Deploy Yaml router initialized successfully")
	}
}

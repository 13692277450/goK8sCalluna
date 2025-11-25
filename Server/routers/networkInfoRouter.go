package routers

import (
	"context"
	"gok8s/serverServices"

	"github.com/gin-gonic/gin"
)

func NetworkInfoRoutersInit(r *gin.Engine) {
	// 使用 NewK8sNetworkCollector 创建并初始化收集器实例
	networkInfoC, err := serverServices.NewK8sNetworkCollector()
	if err != nil {
		panic(err) // 或者可以返回错误，取决于您的错误处理策略
	}
	r.GET("/api/networkinfo", networkInfoHandler(networkInfoC.CollectAllNetworkInfo))
}

// networkInfoHandler 适配器函数，将 gin.Context 转换为 context.Context
func networkInfoHandler(handler func(ctx context.Context) (*serverServices.NetworkInfo, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		networkInfo, err := handler(ctx)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, networkInfo)
	}
}

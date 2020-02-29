package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-example/middleware/jwt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers/api"
	v1 "go-gin-example/routers/api/v1"
	
	_ "go-gin-example/docs"
)

func InitRouter() *gin.Engine {
	//新的无中间件服务
	r := gin.New()

	//使用中间价，日志，panic recover
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//启动模式
	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth", api.GetAuth)

	//路由蓝图
	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新增标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//	获取文章
		apiv1.GET("articles/:id", v1.GetArticle)
		//	添加文章
		apiv1.POST("/articles", v1.AddArticle)
		//	修改文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//	删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}

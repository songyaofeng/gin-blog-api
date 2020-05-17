package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-blog-api/api/v1"
	"net/http"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	//engine.GET("/auth", v1.GetAuth)
	apiV1 := engine.Group("/api/v1")
	apiV1.Use()
	{
		// 标签
		apiV1.POST("/tag/list", v1.GetTags)
		apiV1.POST("/tag/create", v1.TagCreate)
		apiV1.POST("/tag/update", v1.TagUpdate)
		apiV1.POST("/tag/delete", v1.TagDelete)

		// 文章
		apiV1.POST("/article/create", v1.ArticleCreate)
		apiV1.POST("/article/update", v1.ArticleUpdate)
		apiV1.POST("/article/list", v1.ArticleList)
		apiV1.POST("/article/delete", v1.ArticleDelete)

	}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return engine
}

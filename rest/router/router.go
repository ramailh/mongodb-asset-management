package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ramailh/mongodb-asset-management/rest/controller"
)

func NewRouter() *gin.Engine {
	rtr := gin.Default()

	rtr.Use(cors.Default())

	blog := rtr.Group("/menanam/blog")
	{
		blog.GET("/", controller.FindBlog)
		blog.GET("/:id", controller.FindBlogByID)
		blog.POST("/", controller.InsertBlog)
		blog.PUT("/:id", controller.UpdateBlog)
		blog.DELETE("/:id", controller.DeleteBlog)
	}

	return rtr
}

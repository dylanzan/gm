package function

import (
"github.com/gin-gonic/gin"
)

func Handler(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	r.GET("/iqiyi", func(context *gin.Context) {
		targetUrl:=context.Query("t")
		context.Redirect(301,targetUrl)
	})
	r.GET("/tencent",func(context *gin.Context) {
		targetUrl:=context.Query("t")
		context.Redirect(301,targetUrl)
	})
	r.GET("/youku",func(context *gin.Context) {
		targetUrl:=context.Query("t")
		context.Redirect(301,targetUrl)
	})
	r.GET("/mango",func(context *gin.Context) {
		targetUrl:=context.Query("t")
		context.Redirect(301,targetUrl)
	})

}

package api

import (
	"gcron-api/util/request"
	"github.com/gin-gonic/gin"
)

func LoadTestRoute(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		r := request.New(c)
		r.Success(r.Gets())
	})
	r.POST("/test", func(c *gin.Context) {
		r := request.New(c)
		r.Success(r.Posts())
	})
}

package api

import (
	"gcron-api/service/job"
	"gcron-api/util/request"
	"github.com/gin-gonic/gin"
)

func LoadJobRoute(r *gin.Engine) {
	r.POST("/job/add", func(c *gin.Context) {
		job.Add(request.New(c))
	})
}

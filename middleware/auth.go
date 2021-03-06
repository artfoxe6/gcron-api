package middleware

import (
	"github.com/gin-gonic/gin"
	"gcron-api/util/request"
	"gcron-api/util/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := request.New(c)

		t := r.Header("Authorization")
		if t == "" {
			r.Error("认证失败")
			r.C.Abort()
		}
		_, err := token.VerifyJwtToken(t)
		if err != nil {
			r.Error("认证失败")
			r.C.Abort()
		}
		//fmt.Println("middleware before")
		r.C.Next()
		//fmt.Println("middleware after")
	}
}

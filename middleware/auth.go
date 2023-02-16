package middleware

import (
	"github.com/gin-gonic/gin"
	"go-blog/controller"
	"go-blog/pkg/jwt"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			controller.RespError(c, controller.CodeNeedLogin, nil)
			c.Abort()
			return
		}

		// 按空格分割，因为格式为 Bearer xxxxx.xxx.xxx
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.RespError(c, controller.CodeInvalidToken, nil)
			c.Abort()
			return
		}

		// parts[1]是获取的 token，对其解析
		myClaims, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.RespError(c, controller.CodeInvalidToken, nil)
			c.Abort()
			return
		}

		// 将当前的 userID 信息保存到请求的上下文 c 上
		c.Set(controller.CtxUserIDKey, myClaims.UserID)
		c.Next() // 后续处理请求的函数中，可以用过 c.Get(CtxUserIDKey) 来获取当前用户的信息
	}
}

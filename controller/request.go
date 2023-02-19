package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const CtxUserIDKey = "userID"

var ErrUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrUserNotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	return
}

// 获取分页参数
func getPageParam(c *gin.Context) (int64, int64) {
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")

	pageNum, err := strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 1
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		// 默认一页展示10个数据量
		pageSize = 10
	}

	return pageNum, pageSize
}

package controller

import (
	"go-blog/api"
	"go-blog/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAllArticle 获取所有文章
func GetAllArticle(c *gin.Context) {
	// 获取分页参数
	pageNumStr, pageSizeStr := getPageParam(c)

	// 操作数据库
	articleList, err := service.QueryArticleList(pageNumStr, pageSizeStr)
	if err != nil {
		zap.L().Error("service.QueryArticleList failed", zap.Error(err))
		RespError(c, CodeServiceBusy, nil)
		return
	}

	RespSuccess(c, gin.H{
		"articleList": articleList,
	})
}

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	// 绑定参数
	param := new(api.PostArticle)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("CreateArticle with invalid params", zap.Error(err))
		RespError(c, CodeInvalidParams, nil)
		return
	}

	// 操作数据库
	if err := service.CreateArticle(param); err != nil {
		zap.L().Error("service.CreateArticle failed", zap.Error(err))
		RespError(c, CodeServiceBusy, nil)
		return
	}

	RespSuccess(c, nil)
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	articleIDStr := c.Param("id")
	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		zap.L().Error("deleteArticle with invalid params", zap.Error(err))
		RespError(c, CodeInvalidParams, nil)
		return
	}

	if err := service.DeleteArticle(articleID); err != nil {
		zap.L().Error("service.DeleteArticle failed", zap.Error(err))
		RespError(c, CodeServiceBusy, nil)
		return
	}

	RespSuccess(c, nil)
}

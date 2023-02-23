package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"go.uber.org/zap"
)

// GetAllArticle 获取所有文章
func GetAllArticle(c *gin.Context) {
	articleList, err := service.QueryArticleList()
	if err != nil {
		zap.L().Error("service.QueryArticleList failed", zap.Error(err))
		RespError(c, CodeServiceBusy, nil)
		return
	}

	RespSuccess(c, gin.H{
		"articleList": articleList,
	})
}

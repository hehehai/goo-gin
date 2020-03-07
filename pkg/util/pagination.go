package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-example/pkg/setting"
)

//分页，更具请求页返回查询位移值
func GetPage(c *gin.Context) int {
	result := 0
	//分页 请求值获取
	page, _ := com.StrTo(c.Query("page")).Int()

	if page > 0 {
		//查询1，数据库 0
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
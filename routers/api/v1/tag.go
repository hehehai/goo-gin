package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/util"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	//查询参数
	name := c.Query("name")

	//查询map
	maps := make(map[string]interface{})
	//响应数据
	data := make(map[string]interface{})

	//是否为空
	if name != "" {
		maps["name"] = name
	}

	//默认状态
	var state int = -1
	//获取查询参数
	if arg := c.Query("state"); arg != "" {
		//类型转换
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	//响应码
	code := e.SUCCESS

	//查询 数据，总数
	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	//响应 json
	c.JSON(http.StatusOK, gin.H{
		"code": code, // 响应码
		"msg":  e.GetMsg(code), // 响应码 描述信息
		"data": data, // 数据
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	//查询名称
	name:= c.Query("name")
	//查询默认名称，类型转换
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	//创建者
	createdBy := c.Query("created_by")

	//数据验证，错误提示信息
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空") // 不为空
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符") // 文字长度
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1") // int 范围

	//响应码
	code := e.INVALID_PARAMS
	//是否验证通过
	if !valid.HasErrors() {
		//标签是否已存在
		if !models.ExistTagByName(name) {
			//状态码
			code = e.SUCCESS
			//创建标签
			models.AddTag(name, state, createdBy)
		} else {
			//标签名已存在
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	//响应 JSON
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//编辑标签
func EditTag(c *gin.Context) {
	//获取动态路由的变量参数
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modifiedBy"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

//删除标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

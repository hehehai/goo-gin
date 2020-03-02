package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-example/pkg/setting"
	"log"
	"time"
)

//数据
var db *gorm.DB

//基础模型结构
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

//初始化数据库连接
func init() {
	//错误，连接数据库相关配置信息
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	//获取配置文件内的数据库相关配置
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//获取
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DB_NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	//连接数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))

	if err != nil {
		log.Println(err)
	}

	// 配置 表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	// run callback for create, update
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	// run callback for delete
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	//单表
	db.SingularTable(true)
	//日志
	db.LogMode(true)
	//最大等待
	db.DB().SetMaxIdleConns(10)
	//最大开启
	db.DB().SetMaxOpenConns(100)
}

//关闭连接
func CloseDB() {
	defer db.Close()
}

// gorm callback init created_on or modified_on timeStamp
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		timeNow := time.Now().Unix()

		// init created time
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(timeNow)
			}
		}

		//init modified time
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(timeNow)
			}
		}
	}
}

// update callback for modified_on
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// delete callback
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		//判断是否有删除时间
		deletedField, hasDeletedField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedField {
			//	无删除时间，执行软删除
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			//	执行硬删除
			scope.Raw(fmt.Sprintf(
				"UPDATE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

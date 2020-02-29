package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-example/pkg/setting"
	"log"
)

//数据
var db *gorm.DB

//基础模型结构
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
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

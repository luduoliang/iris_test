package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"iris_test/config"
)

const (
	PAGE_LIMIT = 10  //分页每页显示条数，默认值10
)

//单例，开启事务时用这个db(models.db)db.Begin(),db.Rollback(),db.Commit()
var db *gorm.DB

//连接数据库，创建表
func init() {
	//根据配置模式不同，读取数据库配置信息
	runmode := config.GetString("runmode", "dev")
	database_driver := config.GetString(runmode + "." + "database_driver", "mysql")
	database_connect_string := config.GetString(runmode + "." + "database_connect_string", "")

	var err error
	//打开数据库链接，返回db实例
	db, err = gorm.Open(database_driver, database_connect_string)
	if err != nil {
		log.Println("错误：", err)
	}
	//取消建表时表名自动添加“s”
	db.SingularTable(true)
	//创建用户表
	if !db.HasTable(&Admin{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台用户表'").CreateTable(&Admin{}).Error; err != nil {
			log.Println(err)
		}
	}
	//创建资讯分类表
	if !db.HasTable(&Category{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资讯分类表'").CreateTable(&Category{}).Error; err != nil {
			log.Println(err)
		}
	}
	//创建资讯表
	if !db.HasTable(&News{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资讯表'").CreateTable(&News{}).Error; err != nil {
			log.Println(err)
		}
	}
	log.Println("数据库已连接")
}
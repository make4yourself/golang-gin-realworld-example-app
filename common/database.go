package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type Database struct {
	*gorm.DB
}

// 全局 DB 变量，是程序与数据库的连接管理状态
var DB *gorm.DB

// 建立与数据库之间的连接，并将连接状态用一个全局变量 DB 保存
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./common/DBdata/gorm.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

// 这个函数会创建一个临时的数据库用于测试样例
func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("sqlite3", "./../gorm_test.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

// 删除为测试样例创建的临时数据库
func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../gorm_test.db")
	return err
}

// 使用这个函数从全局变量中或取已经获得的数据库连接
func GetDB() *gorm.DB {
	return DB
}

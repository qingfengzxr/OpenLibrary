/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:48
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 08:33:10
 * @FilePath: \ol\internal\dbgorm\db.go
 */
package dbgorm

import (
	"fmt"
	"os"

	"ol/config/impl"
	"ol/internal/dbgorm/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"Gorm database": "",
	})
)

// InitGormDB InitGormDB
func InitGormDB(driveName string, cfg impl.DatabaseConfig) (*gorm.DB, error) {
	log.Infoln("InitGormDB")
	if driveName == "" {
		return nil, os.ErrInvalid
	}

	// 为了处理time.Time，需要包括parseTime作为参数
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", cfg.User, cfg.Passwd, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset)
	db, err := gorm.Open(driveName, dbDSN)
	if err != nil {
		return nil, err
	}

	// 全局禁用表名复数
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 使用如下方式设置日志输出级别以及改变日志输出地方
	// db.LogMode(true)
	db.SetLogger(log.Logger)

	err = checkDatabaseAndTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func checkDatabaseAndTables(dbMgr *gorm.DB) error {
	log.Infoln("checkDatabaseAndTables")
	if dbMgr == nil {
		return os.ErrInvalid
	}

	// 用户关系表
	if !dbMgr.HasTable(&model.Book{}) {
		// 通过AutoMigrate函数可以快速建表，如果表已经存在不会重复创建。
		// dbMgr.AutoMigrate(&model.Book{})
		dbMgr.CreateTable(&model.Book{})
	}

	return nil
}

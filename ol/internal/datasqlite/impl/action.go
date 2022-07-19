/*
 * @Author: tj
 * @Date: 2022-07-19 18:58:32
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 21:38:28
 * @FilePath: \OpenLibrary\ol\internal\datasqlite\impl\action.go
 */
package impl

import (
	"ol/internal/datasqlite/handler"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"Action": "",
	})
)

//Action 函数对象
type Action struct {
}

func NewAction() *Action {
	return &Action{}
}

//CreateTable 建表
func (act *Action) CreateTable(tableName string) error {
	db, err := handler.NewSqliteDB()
	if err != nil {
		log.Errorln("CreateTable NewSqliteDB error:", err)
		return err
	}

	return db.CreateTable(tableName)
}

//DeleleTable 删表
func (act *Action) DeleleTable(tableName string) error {
	db, err := handler.NewSqliteDB()
	if err != nil {
		return err
	}

	return db.DeleteTable(tableName)
}

//InsertData 插入数据
func (act *Action) InsertData(tableName string, destDataKeys []string, destDataValues []interface{}) error {
	db, err := handler.NewSqliteDB()
	if err != nil {
		return err
	}

	return db.InsertData(tableName, destDataKeys, destDataValues)
}

//ReadData 读取数据
func (act *Action) ReadData(tableName string, selectKeys []string, selectValues []interface{}, destDataKeys []string) ([][]string, error) {
	db, err := handler.NewSqliteDB()
	if err != nil {
		return nil, err
	}

	return db.ReadData(tableName, selectKeys, selectValues, destDataKeys)
}

//UpData 修改数据
func (act *Action) UpData(tableName string, selectKeys []string, selectValues []interface{}, destDataKeys []string, destDataValues []interface{}) error {
	db, err := handler.NewSqliteDB()
	if err != nil {
		return err
	}

	return db.UpData(tableName, selectKeys, selectValues, destDataKeys, destDataValues)
}

//DeleteData 删除数据
func (act *Action) DeleteData(tableName string, selectKeys []string, selectValues []interface{}) error {
	db, err := handler.NewSqliteDB()
	if err != nil {
		return err
	}

	return db.DeleteData(tableName, selectKeys, selectValues)
}

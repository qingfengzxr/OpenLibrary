package handler

import (
	"database/sql"
	"errors"
	"fmt"

	"ol/internal/datasqlite/data"
	dalsql "ol/internal/sqlite"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"RecordHandle": "",
	})
)

//DBsqlite database of sqlite
type DBsqlite struct {
	DB *sql.DB
}

//NewSqliteDB 初始化DB并返回
func NewSqliteDB() (*DBsqlite, error) {
	db := dalsql.GetInstance().RawHandler
	if db == nil {
		return nil, errors.New("sql db is nil")
	}

	newDB := &DBsqlite{
		DB: db,
	}

	return newDB, nil
}

//CreateTable 建表
func (db *DBsqlite) CreateTable(tableName string) error {
	createTable, ok := data.TableCreate[tableName]
	if !ok {
		return errors.New("the table Unkown")
	}

	result, err := db.DB.Exec(createTable)
	if err != nil {
		return err
	}
	log.Infof("CreateTable : %#v", result)

	return nil
}

//DeleteTable 删表
func (db *DBsqlite) DeleteTable(tableName string) error {
	deleteTabel := fmt.Sprintf("DROP TABLE %s;", tableName)
	result, err := db.DB.Exec(deleteTabel)
	if err != nil {
		return err
	}
	log.Infof("Delete Table : %#v", result)

	return nil
}

//InsertData 插入数据
func (db *DBsqlite) InsertData(tableName string, destDataKeys []string, destDataValues []interface{}) error {
	var parmKey, parmValueMark string
	for _, key := range destDataKeys {
		parmKey += fmt.Sprintf("%s,", key)
		parmValueMark += "?,"
	}

	insertDataSQL := fmt.Sprintf("INSERT INTO %s(%s) values(%s)", tableName, parmKey[:len(parmKey)-1], parmValueMark[:len(parmValueMark)-1])
	log.Debug("insertDataSQL: ", insertDataSQL)

	stmt, err := db.DB.Prepare(insertDataSQL)
	if err != nil {
		return err
	}

	log.Debugf("destDataValues: %#v", destDataValues)

	result, err := stmt.Exec(destDataValues...)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	log.Debugf("inserted id is %d", lastID)

	return nil
}

//ReadData 读取数据
func (db *DBsqlite) ReadData(tableName string, selectKeys []string, selectValues []interface{}, destDataKeys []string) ([][]string, error) {
	if tableName == "" {
		return nil, errors.New("ReadData tableName is empty")
	}

	selectParm := "init"
	if selectKeys != nil {
		if len(selectKeys) != len(selectValues) {
			return nil, errors.New("select data not match in request")
		}

		if len(selectKeys) != 0 {
			selectParm = " WHERE"
			for i := range selectKeys {
				selectParm += fmt.Sprintf(" %s=? and", selectKeys[i])
			}
		}
	}

	destParm := "* "
	if destDataKeys != nil {
		destParm = ""
		for _, key := range destDataKeys {
			destParm += fmt.Sprintf("%s,", key)
		}
	}
	readDataSQL := fmt.Sprintf("SELECT %s FROM %s%s", destParm[:len(destParm)-1], tableName, selectParm[:len(selectParm)-4])
	log.Debug("readDataSQL: ", readDataSQL)

	stmt, err := db.DB.Prepare(readDataSQL)
	if err != nil {
		return nil, err
	}

	log.Debugf("selectValues: %#v", selectValues)

	rows, err := stmt.Query(selectValues...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dstDataRows := make([][]string, 0)

	for rows.Next() {
		dstDataRow, err := getDataFromRows(rows)
		if err != nil {
			return nil, err
		}

		dstDataRows = append(dstDataRows, dstDataRow)
	}

	log.Debugf("dstDataRows: %#v", dstDataRows)

	return dstDataRows, nil
}

//UpData 修改数据
func (db *DBsqlite) UpData(tableName string, selectKeys []string, selectValues []interface{}, destDataKeys []string, destDataValues []interface{}) error {
	log.Debugf("tableName : %s", tableName)
	log.Debugf("selectKeys : %#v", selectKeys)
	log.Debugf("selectValues : %#v", selectValues)
	log.Debugf("destDataKeys : %#v", destDataKeys)
	log.Debugf("destDataValues : %#v", destDataValues)

	if tableName == "" {
		return errors.New("UpData tableName is empty")
	}

	selectParm := "init"
	if selectKeys != nil {
		if len(selectKeys) != len(selectValues) {
			return errors.New("select data not match in request")
		}

		if len(selectKeys) != 0 {
			selectParm = " WHERE"
			for i := range selectKeys {
				selectParm += fmt.Sprintf(" %s=? and", selectKeys[i])
			}
		}
	}

	var destParm string
	if destDataKeys == nil {
		return errors.New("dest data is nil")
	}

	if len(destDataKeys) == 0 {
		return errors.New("dest data is empty")
	}

	if len(destDataKeys) != len(destDataValues) {
		return errors.New("dest data not match in request")
	}

	for i := range destDataKeys {
		destParm += fmt.Sprintf(" %s=?,", destDataKeys[i])
	}

	upDataSQL := fmt.Sprintf("UPDATE %s SET%s%s", tableName, destParm[:len(destParm)-1], selectParm[:len(selectParm)-4])
	log.Debug("upDataSQL: ", upDataSQL)

	stmt, err := db.DB.Prepare(upDataSQL)
	if err != nil {
		return err
	}

	sumValues := make([]interface{}, 0)

	for _, destValue := range destDataValues {
		sumValues = append(sumValues, destValue)
	}

	for _, selectValue := range selectValues {
		sumValues = append(sumValues, selectValue)
	}

	log.Debugf("sumValues: %#v", sumValues)

	result, err := stmt.Exec(sumValues...)
	if err != nil {
		return err
	}

	affectNum, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Debugf("update affect rows is %d", affectNum)

	return nil
}

//DeleteData 删除数据
func (db *DBsqlite) DeleteData(tableName string, selectKeys []string, selectValues []interface{}) error {

	if tableName == "" {
		return errors.New("tableName is empty")
	}

	selectParm := "init"
	if selectKeys != nil {
		if len(selectKeys) != len(selectValues) {
			return errors.New("select data not match in request")
		}

		if len(selectKeys) != 0 {
			selectParm = " WHERE"
			for i := range selectKeys {
				selectParm += fmt.Sprintf(" %s=? and", selectKeys[i])
			}
		}
	}

	deleteDataSQL := fmt.Sprintf("DELETE FROM %s%s", tableName, selectParm[:len(selectParm)-4])
	log.Debug("deleteDataSQL: ", deleteDataSQL)

	stmt, err := db.DB.Prepare(deleteDataSQL)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(selectValues...)
	if err != nil {
		return err
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Debugf("delete affect rows is %d", affectNum)

	return nil
}

func getDataFromRows(rows *sql.Rows) ([]string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	err = rows.Scan(scanArgs...)
	if err != nil {
		return nil, err
	}

	dstDataRow := make([]string, len(values))
	for i := range values {
		dstDataRow[i] = string(values[i])
	}

	return dstDataRow, nil
}

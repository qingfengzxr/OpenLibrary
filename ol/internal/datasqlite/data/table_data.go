/*
 * @Author: tj
 * @Date: 2022-07-19 18:58:32
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 19:44:46
 * @FilePath: \ol\internal\datasqlite\data\table_data.go
 */
package data

const (
	//Book 表：书内容表
	Book = "Book"
	//BookCreate 构表语句
	BookCreate = `
    CREATE TABLE IF NOT EXISTS Book(
        Cid         VARCHAR(255)    NOT NULL,
        Creator     VARCHAR(255)    NOT NULL,
        Chapter     int32           NOT NULL,
        Section     int32           NOT NULL,
        Page        int32           NOT NULL,
        Content     VARCHAR(255)    NOT NULL,
        CreateAt    TEXT            NOT NULL,
        UpdateAt    TEXT
    );
    `
)

//TableCreate 构表映射
var TableCreate = map[string]string{Book: BookCreate}

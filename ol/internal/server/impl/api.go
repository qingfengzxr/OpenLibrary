/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:38
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 20:08:34
 * @FilePath: \ol\internal\server\impl\api.go
 */
package impl

import (
	"encoding/base64"
	"os"
	"time"

	sqliteData "ol/internal/datasqlite/data"
	"ol/internal/dbgorm/model"
	"ol/internal/server/data"
)

const (
	short = "2006-01-02 15:04:05"
)

func (m *Server) SaveContent(req *data.SaveRequest) (*data.SaveResponse, error) {
	if req == nil || req.Cid == "" || req.Creator == "" || req.Page <= 0 || len(req.Content) == 0 {
		return nil, os.ErrInvalid
	}

	// 卷:可有可无，默认1
	if req.Chapter <= 0 {
		req.Chapter = 1
	}

	if m.dbMgr != nil {
		err := m.dbMgr.Create(&model.Book{
			Cid:      req.Cid,
			Creator:  req.Creator,
			Chapter:  req.Chapter,
			Section:  req.Section,
			Page:     req.Page,
			Content:  []byte(base64.StdEncoding.EncodeToString(req.Content)),
			CreateAt: time.Now().Local(),
			UpdateAt: time.Now().Local(),
			DeleteAt: time.Now().Local(),
		}).Error
		if err != nil {
			log.Errorln("Create book info error:", err)
			return nil, err
		}
	} else {
		createTime := time.Now().String()
		updateTime := time.Now().String()

		destKeys := []string{"Cid", "Creator", "Chapter", "Section", "Page", "Content", "CreateAt", "UpdateAt"}
		destValues := []interface{}{req.Cid, req.Creator, req.Chapter, req.Section, req.Page, base64.StdEncoding.EncodeToString(req.Content), createTime[:19], updateTime[:19]}
		err := m.sqliteDb.InsertData(sqliteData.Book, destKeys, destValues)
		if err != nil {
			log.Errorln("Create book info error:", err)
			return nil, err
		}
	}

	return nil, nil
}

func (m *Server) GetContent(req *data.GetRequest) (*data.GetResponse, error) {
	if req == nil || req.Cid == "" {
		return nil, os.ErrInvalid
	}

	if m.dbMgr != nil {
		book := &model.Book{}
		m.dbMgr.Where(model.Book{Cid: req.Cid}).First(&book)

		if book.Cid != "" {
			content, err := base64.StdEncoding.DecodeString(string(book.Content))
			if err != nil {
				return nil, err
			}

			return &data.GetResponse{Content: content}, nil
		}
	} else {
		selectKeys := []string{"Cid"}
		selectValue := []interface{}{req.Cid}

		dstDataRows, err := m.sqliteDb.ReadData(sqliteData.Book, selectKeys, selectValue, nil)
		if err != nil {
			return nil, err
		}

		if len(dstDataRows) != 0 {
			content, err := base64.StdEncoding.DecodeString(dstDataRows[0][5])
			if err != nil {
				return nil, err
			}
			return &data.GetResponse{Content: content}, nil
		}

	}

	return nil, nil
}

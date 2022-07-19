/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:38
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:51:04
 * @FilePath: \ol\internal\server\impl\api.go
 */
package impl

import (
	"encoding/base64"
	"os"
	"time"

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

	return nil, nil
}

func (m *Server) GetContent(req *data.GetRequest) (*data.GetResponse, error) {
	if req == nil || req.Cid == "" {
		return nil, os.ErrInvalid
	}

	book := &model.Book{}
	m.dbMgr.Where(model.Book{Cid: req.Cid}).First(&book)

	if book.Cid != "" {
		content, err := base64.StdEncoding.DecodeString(string(book.Content))
		if err != nil {
			return nil, err
		}

		return &data.GetResponse{Content: content}, nil
	}

	return nil, nil
}

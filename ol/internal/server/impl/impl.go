/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:38
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:51:27
 * @FilePath: \ol\internal\server\impl\impl.go
 */
package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"Server": "",
	})
)

// Server server
type Server struct {
	dbMgr *gorm.DB
}

// NewServer return Server
func NewServer(dbMgr *gorm.DB) *Server {
	m := &Server{
		dbMgr: dbMgr,
	}

	return m
}

// Start Start
func (m *Server) Start() error {
	log.Infoln("Server Start success")
	return nil
}

// Stop Stop
func (m *Server) Stop() error {
	m.dbMgr.Close()

	return nil
}

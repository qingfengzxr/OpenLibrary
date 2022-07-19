/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:58
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:44:32
 * @FilePath: \ol\api\httpserver\impl\impl.go
 */
package impl

import (
	"fmt"
	"io"
	"os"
	"time"

	"ol/internal/server"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"GinHTTPServer": "",
	})
)

// GinHTTPServer GinHTTPServer
type GinHTTPServer struct {
	e                 *gin.Engine
	addr              string
	port              string
	eventListenServer server.EventListener
}

// NewGinHTTPServer NewGinHTTPServer
func NewGinHTTPServer(address, port string) *GinHTTPServer {
	http := &GinHTTPServer{
		e:    gin.Default(),
		addr: address,
		port: port,
	}
	return http
}

// Start Start
func (g *GinHTTPServer) Start() error {
	log.Info("GinHTTPServer Start")

	// 输出日志到文件
	gin.DefaultWriter = io.MultiWriter(log.Writer())
	g.e.Use(gin.Logger())

	g.e.Use(Cors())
	g.e.Use(TimeOut(time.Second * 3))

	pprof.Register(g.e)

	g.setRouteGroup()

	url := fmt.Sprintf("%s:%s", g.addr, g.port)
	err := g.e.Run(url)
	if err != nil {
		return err
	}

	return nil
}

// SetEventListenServer SetEventListenServer
func (g *GinHTTPServer) SetEventListenServer(svr server.EventListener) error {
	if svr == nil {
		return os.ErrInvalid
	}

	g.eventListenServer = svr
	return nil
}

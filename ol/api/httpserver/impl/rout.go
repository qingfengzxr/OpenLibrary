/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:58
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:45:07
 * @FilePath: \ol\api\httpserver\impl\rout.go
 */
package impl

import (
	"net/http"

	"ol/internal/server/data"

	"github.com/gin-gonic/gin"
)

func (g *GinHTTPServer) setRouteGroup() {
	log.Info("setRouteGroup")
	group := g.e.Group("api/")

	group.POST("book/save", g.save)
	group.POST("book/get", g.get)
}

func (g *GinHTTPServer) save(c *gin.Context) {
	ret := &result{}
	defer c.JSON(http.StatusOK, ret)

	req := &data.SaveRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Errorln("save BindJSON error:", err)
		ret.Code = 1
		ret.Msg = err.Error()
		return
	}

	response, err := g.eventListenServer.SaveContent(req)
	if err != nil {
		ret.Code = 1
		log.Errorln("save SaveContent error:", err)
		ret.Msg = err.Error()
		return
	}

	if response != nil {
		ret.Data = response
	}
}

func (g *GinHTTPServer) get(c *gin.Context) {
	ret := &result{}
	defer c.JSON(http.StatusOK, ret)

	req := &data.GetRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Errorln("get BindJSON error:", err)
		ret.Code = 1
		ret.Msg = err.Error()
		return
	}

	response, err := g.eventListenServer.GetContent(req)
	if err != nil {
		ret.Code = 1
		log.Errorln("get GetContent error:", err)
		ret.Msg = err.Error()
		return
	}

	if response != nil {
		ret.Data = response
	}
}

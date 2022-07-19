/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:38
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:51:32
 * @FilePath: \ol\internal\server\interface.go
 */
package server

import "ol/internal/server/data"

type EventListener interface {
	Start() error
	Stop() error

	SaveContent(req *data.SaveRequest) (*data.SaveResponse, error)
	GetContent(req *data.GetRequest) (*data.GetResponse, error)
}

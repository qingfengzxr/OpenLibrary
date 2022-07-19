/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:58
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:45:52
 * @FilePath: \ol\api\httpserver\httpserver.go
 */
package httpserver

import (
	"ol/internal/server"
)

// GinServer GinServer
type GinServer interface {
	Start() error

	SetEventListenServer(svr server.EventListener) error
}

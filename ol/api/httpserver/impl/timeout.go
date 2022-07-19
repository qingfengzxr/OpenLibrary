/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:58
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:44:32
 * @FilePath: \ol\api\httpserver\impl\impl.go
 */
package impl

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeOut(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			cancel()
		}()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

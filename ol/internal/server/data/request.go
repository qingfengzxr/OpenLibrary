/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:48
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:24:41
 * @FilePath: \ol\internal\server\data\request.go
 */
package data

type SaveRequest struct {
	Cid     string `json:"cid" binding:"required"`
	Creator string `json:"creator" binding:"required"` // 创建者
	Chapter int    `json:"chapter"`                    // 卷:可有可无，默认1
	Section int    `json:"section" binding:"required"` // 章节
	Page    int    `json:"page" binding:"required"`    // 页数
	Content []byte `json:"content" binding:"required"` // 内容
}

type GetRequest struct {
	Cid string `json:"cid" binding:"required"`
}

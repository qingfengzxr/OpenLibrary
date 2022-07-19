/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:58
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:44:22
 * @FilePath: \ol\api\httpserver\impl\const.go
 */
package impl

const (
	ErrRepeatSubmit = "请勿重复提交申请"
	ErrTokenInvalid = "无效的token"

	AvoidRepeatedSubmissionsTimeOut = 3
)

type result struct {
	Code int
	Data interface{}
	Msg  string
}

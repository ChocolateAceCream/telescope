/*
* @fileName shared.go
* @author Di Sheng
* @date 2022/10/18 08:49:42
* @description: wrap out http response
 */

package response

import (
	"net/http"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"error_code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Paging struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

const (
	ERROR        = 1
	SUCCESS      = 0
	UNAUTHORIZED = 401
)

func translate(c *gin.Context, raw string) (translated string) {
	mapper := singleton.Translation
	lang := c.GetHeader("X-Language")
	dictionary, ok := mapper[lang]
	if !ok {
		dictionary = mapper["en"]
	}
	translated, ok = dictionary[raw]
	if !ok {
		return raw
	}
	return
}

func ResponseGenerator(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		translate(c, msg),
	})
}

func OkWithFullDetails(c *gin.Context, data interface{}, msg string) {
	ResponseGenerator(c, SUCCESS, data, msg)
}

func OkWithMessage(c *gin.Context, msg string) {
	ResponseGenerator(c, SUCCESS, map[string]interface{}{}, msg)
}

func FailWithMessage(c *gin.Context, msg string) {
	ResponseGenerator(c, ERROR, map[string]interface{}{}, msg)
}

func FailWithFullDetails(c *gin.Context, data interface{}, msg string) {
	ResponseGenerator(c, ERROR, data, msg)
}

func FailWithUnauthorized(c *gin.Context, msg string) {
	ResponseGenerator(c, UNAUTHORIZED, map[string]interface{}{}, msg)

}

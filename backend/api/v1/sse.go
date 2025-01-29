package apiV1

import (
	"fmt"

	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/gin-gonic/gin"
)

type SSEApi struct{}

func (s *SSEApi) Subscriber(c *gin.Context) {
	fmt.Println("header check", c.Writer.Header())
	err := sseService.Subscribe(c)
	if err != nil {
		response.FailWithMessage(c, "fail to subscribe")
	}
}

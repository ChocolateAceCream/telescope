package service

import (
	"io"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/lib"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SSEService struct{}

func (s *SSEService) Subscribe(c *gin.Context) (err error) {
	user, err := utils.GetValueFromSessionByKey[db.AUser](c, "user")
	if err != nil {
		singleton.Logger.Error("fail to get user from session", zap.Error(err))
		return
	}
	uid := user.ID
	clientChan := lib.GetActiveSSE(uid)
	go func() {
		<-c.Writer.CloseNotify()
		singleton.Logger.Info("Client disconnected")
		lib.DeactivateSSE(uid)
	}()
	c.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		if msg, ok := <-clientChan; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})

	return
}

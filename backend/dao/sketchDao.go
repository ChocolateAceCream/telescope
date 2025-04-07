package dao

import (
	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SketchDao struct{}

func (s *SketchDao) CreateSketch(c *gin.Context, payload db.NewSketchParams) (sketch db.Sketch, err error) {
	sketch, err = singleton.Query.NewSketch(c, payload)
	if err != nil {
		singleton.Logger.Error("CreateSketch failed", zap.Error(err))
	}
	return
}

func (s *SketchDao) GetSketchesByProjectID(c *gin.Context, projectID int) (sketches []db.Sketch, err error) {
	sketches, err = singleton.Query.GetSketchesByProjectID(c, int32(projectID))
	if err != nil {
		singleton.Logger.Error("GetSketchByProjectID failed", zap.Error(err))
	}
	return
}

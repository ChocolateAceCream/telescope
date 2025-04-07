package apiV1

import (
	"fmt"

	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
)

type SketchApi struct{}

func (s *SketchApi) UploadSketch(c *gin.Context) {
	user, err := utils.GetSessionUser(c)
	if err != nil {
		response.FailWithMessage(c, "error.failed.operation")
		return
	}
	fmt.Println("------user------", user)
	var req request.UploadSketchRequest
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("------err------", err)
		response.FailWithMessage(c, "error.missing.params")
		return
	}
	err = sketchService.UploadSketch(c, user, req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithMessage(c, "success")
}

func (s *SketchApi) GetSketchList(c *gin.Context)   {}
func (s *SketchApi) GetSketchDetail(c *gin.Context) {}
func (s *SketchApi) UpdateSketch(c *gin.Context)    {}
func (s *SketchApi) DeleteSketch(c *gin.Context)    {}

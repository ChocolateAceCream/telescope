package apiV1

import (
	"strconv"

	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProjectApi struct{}

func (p *ProjectApi) GetProjectList(c *gin.Context) {
	var req request.ProjectListRequest
	if err := c.ShouldBind(&req); err != nil {
		singleton.Logger.Error("Failed to bind JSON", zap.Error(err))
		response.FailWithMessage(c, "error.missing.params")
		return
	}
	req.Params.ApplyDefaultsAndValidate()
	resp, err := projectService.GetProjectList(c, req.Params)
	if err != nil {
		singleton.Logger.Error("Failed to get project list", zap.Error(err))
		response.FailWithMessage(c, "error.failed.operation")
		return
	}
	response.OkWithFullDetails(c, resp, "success")
}

func (p *ProjectApi) GetProjectDetails(c *gin.Context) {
	id := c.Param("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		singleton.Logger.Error("Failed to convert project ID", zap.Error(err))
		response.FailWithMessage(c, "error.missing.params")
		return
	}
	resp, err := projectService.GetProjectDetails(c, projectID)
	if err != nil {
		singleton.Logger.Error("Failed to get project details", zap.Error(err))
		response.FailWithMessage(c, "error.failed.operation")
		return
	}
	response.OkWithFullDetails(c, resp, "success")
}

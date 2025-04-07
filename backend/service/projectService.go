package service

import (
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/gin-gonic/gin"
)

type ProjectService struct{}

func (p *ProjectService) GetProjectList(c *gin.Context, req request.ProjectListRequestParam) (resp response.ProjectListResponse, err error) {
	resp.Projects, err = projectDao.GetProjectList(c, req)
	if err != nil {
		return
	}
	resp.Total, err = projectDao.GetTotalProjectCount(c)
	return
}

func (p *ProjectService) GetProjectDetails(c *gin.Context, projectID int) (resp response.ProjectDetailsResponse, err error) {
	project, err := projectDao.GetProjectByID(c, projectID)
	if err != nil {
		return
	}
	s, err := sketchDao.GetSketchesByProjectID(c, projectID)
	resp = response.ProjectDetailsResponse{
		Project: response.Project{
			ID:          int(project.ID),
			ProjectName: project.ProjectName,
			Comment:     project.Comment.String,
			Address:     project.Address.String,
			UpdatedAt:   project.UpdatedAt.Time,
			Status:      project.Status.String,
		},
		Sketches: make([]response.Sketch, len(s)),
	}
	for i, sketch := range s {
		resp.Sketches[i] = response.Sketch{
			ID:           int(sketch.ID),
			ProjectName:  project.ProjectName,
			ProjectID:    int(project.ID),
			UpdatedAt:    sketch.UpdatedAt.Time,
			FullImageUrl: sketch.FullImageUrl,
			UploaderID:   int(sketch.UploaderID),
		}
	}
	return
}

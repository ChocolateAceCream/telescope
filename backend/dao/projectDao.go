package dao

import (
	"fmt"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProjectDao struct{}

func (p *ProjectDao) GetProjectByName(c *gin.Context, name string) (project db.Project, err error) {
	project, err = singleton.Query.GetProjectByName(c, name)
	if err != nil {
		singleton.Logger.Error("GetProjectByName failed", zap.Error(err))
	}
	return
}

func (p *ProjectDao) CreateProject(c *gin.Context, payload db.NewProjectParams) (project db.Project, err error) {
	project, err = singleton.Query.NewProject(c, payload)
	if err != nil {
		singleton.Logger.Error("CreateProject failed", zap.Error(err))
	}
	return
}

func (p *ProjectDao) GetProjectList(c *gin.Context, payload request.ProjectListRequestParam) (list []response.Project, err error) {
	query := fmt.Sprintf(
		`
			SELECT id, project_name, comment, status, updated_at, address
			FROM project
			ORDER BY %s %s
			LIMIT $1 OFFSET $2
		`, payload.OrderBy, payload.SortBy)
	rows, err := singleton.DB.Query(c, query, payload.PageSize, (payload.PageNumber-1)*payload.PageSize)
	if err != nil {
		singleton.Logger.Error("GetProjectList failed", zap.Error(err))
		return
	}
	for rows.Next() {
		var p db.Project
		err = rows.Scan(&p.ID, &p.ProjectName, &p.Comment, &p.Status, &p.UpdatedAt, &p.Address)
		if err != nil {
			singleton.Logger.Error("GetProjectList failed", zap.Error(err))
			return
		}
		r := response.Project{
			ID:          int(p.ID),
			ProjectName: p.ProjectName,
			Comment:     p.Comment.String,
			Address:     p.Address.String,
			UpdatedAt:   p.UpdatedAt.Time,
			Status:      p.Status.String,
		}
		list = append(list, r)
	}
	if err = rows.Err(); err != nil {
		singleton.Logger.Error("GetProjectList failed", zap.Error(err))
		return
	}
	return
}

func (p *ProjectDao) GetTotalProjectCount(c *gin.Context) (total int, err error) {
	t, err := singleton.Query.GetTotalProjectCount(c)
	if err != nil {
		singleton.Logger.Error("UpdateProject failed", zap.Error(err))
		return
	}
	total = int(t)
	return
}

func (p *ProjectDao) GetProjectByID(c *gin.Context, projectID int) (project db.Project, err error) {
	project, err = singleton.Query.GetProjectByID(c, int32(projectID))
	if err != nil {
		singleton.Logger.Error("GetProjectDeatilsByID failed", zap.Error(err))
	}
	return
}

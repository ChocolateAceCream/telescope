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

func (p *ProjectDao) UpdateProject(c *gin.Context, payload db.UpdateProjectParams) (project db.Project, err error) {
	project, err = singleton.Query.UpdateProject(c, payload)
	if err != nil {
		singleton.Logger.Error("UpdateProject failed", zap.Error(err))
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
	baseQuery := `
		SELECT id, project_name, comment, status, updated_at, address
		FROM project
	`
	var args []interface{}
	whereClause := ""
	argIndex := 1
	// Add blur search if needed
	if payload.ProjectName != "" {
		whereClause = fmt.Sprintf("WHERE project_name ILIKE $%d", argIndex)
		args = append(args, "%"+payload.ProjectName+"%")
		argIndex++
	}
	// Add LIMIT and OFFSET
	limitClause := fmt.Sprintf("ORDER BY %s %s LIMIT $%d OFFSET $%d", payload.OrderBy, payload.SortBy, argIndex, argIndex+1)
	args = append(args, payload.PageSize, (payload.PageNumber-1)*payload.PageSize)

	// Final query
	query := baseQuery + whereClause + "\n" + limitClause
	rows, err := singleton.DB.Query(c, query, args...)
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

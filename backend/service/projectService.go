package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectService struct{}

func (s *ProjectService) UpdateProject(c *gin.Context, user db.User, projectID int, req request.UpdateProjectRequest) (err error) {
	project, err := projectDao.UpdateProject(c, db.UpdateProjectParams{
		ID:          int32(projectID),
		ProjectName: req.Project,
		Comment:     pgtype.Text{String: req.Comment, Valid: req.Comment != ""},
		Address:     pgtype.Text{String: req.Address, Valid: req.Address != ""},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// project not found, create a new project
			payload := db.NewProjectParams{
				ProjectName: req.Project,
				Comment:     pgtype.Text{String: req.Comment, Valid: req.Comment != ""},
				Creator:     user.ID,
				Status:      pgtype.Text{String: "active", Valid: true},
				Address:     pgtype.Text{String: req.Address, Valid: req.Address != ""},
			}
			project, err = projectDao.CreateProject(c, payload)
			if err != nil {
				return
			}
		} else {
			return
		}
	}

	for _, file := range req.Files {
		f, _ := file.Open()
		defer f.Close()
		_, err = singleton.AWS.S3.PutObject(c, &s3.PutObjectInput{
			Bucket: aws.String(singleton.Config.AWS.S3.SketchPublicBucket),
			Key:    aws.String(singleton.Config.AWS.S3.SketchFolder + "/" + project.ProjectName + "/" + file.Filename),
			Body:   f,
		})
		if err == nil {
			url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/%s/%s", singleton.Config.AWS.S3.SketchPublicBucket, singleton.Config.AWS.Region, singleton.Config.AWS.S3.SketchFolder, project.ProjectName, file.Filename)
			url = strings.ReplaceAll(url, " ", "+")
			_, err = sketchDao.CreateSketch(c, db.NewSketchParams{
				ProjectName:  project.ProjectName,
				ProjectID:    project.ID,
				UploaderID:   user.ID,
				FullImageUrl: url,
			})
		}
	}
	return
}

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

package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type SketchService struct{}

func (s *SketchService) UploadSketch(c *gin.Context, user db.User, req request.UploadSketchRequest) (err error) {
	fmt.Println("------user------", user)
	fmt.Println("------req------", req)
	project, err := projectDao.GetProjectByName(c, req.Project)
	if errors.Is(err, sql.ErrNoRows) {
		// project not found, create a new project
		payload := db.NewProjectParams{
			ProjectName: req.Project,
			Comment:     pgtype.Text{String: req.Comment, Valid: false},
			Creator:     user.ID,
			Status:      pgtype.Text{String: "active", Valid: true},
			Address:     pgtype.Text{String: req.Address, Valid: false},
		}
		project, err = projectDao.CreateProject(c, payload)
		if err != nil {
			return
		}
	}
	fmt.Println("------project------", project)
	for _, file := range req.Files {
		f, _ := file.Open()
		defer f.Close()
		_, err = singleton.AWS.S3.PutObject(c, &s3.PutObjectInput{
			Bucket: aws.String(singleton.Config.AWS.S3.SketchPublicBucket),
			Key:    aws.String(singleton.Config.AWS.S3.SketchFolder + "/" + project.ProjectName + "/" + file.Filename),
			Body:   f,
		})
		if err == nil {
			url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/%s/%s", singleton.Config.AWS.S3.SketchPublicBucket, singleton.Config.AWS.Region, singleton.Config.AWS.S3.SketchFolder, url.PathEscape(project.ProjectName), file.Filename)
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

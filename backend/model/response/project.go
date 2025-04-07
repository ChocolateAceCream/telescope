package response

import (
	"time"
)

type ProjectListResponse struct {
	Total    int       `json:"total"`
	Projects []Project `json:"projects"`
}

type Project struct {
	ID          int       `json:"id" db:"id"`
	ProjectName string    `json:"project_name" db:"project_name"`
	Comment     string    `json:"comment" db:"comment"`
	Address     string    `json:"address" db:"address"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Status      string    `json:"status" db:"status"`
}

type ProjectDetailsResponse struct {
	Project  Project  `json:"project"`
	Sketches []Sketch `json:"sketches"`
}

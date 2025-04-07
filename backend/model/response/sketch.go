package response

import "time"

type Sketch struct {
	ID           int       `json:"id" db:"id"`
	ProjectName  string    `json:"project_name" db:"project_name"`
	ProjectID    int       `json:"project_id" db:"project_id"`
	UploaderID   int       `json:"uploader_id" db:"uploader_id"`
	FullImageUrl string    `json:"full_image_url" db:"full_image_url"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

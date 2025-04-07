package request

import (
	"mime/multipart"
)

type UploadSketchRequest struct {
	Project string                  `form:"project" binding:"required"`
	Comment string                  `form:"comment"`
	Address string                  `form:"address"`
	Files   []*multipart.FileHeader `form:"files" binding:"required"`
}

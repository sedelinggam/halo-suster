package request

import "mime/multipart"

type ImageRequest struct {
	File *multipart.FileHeader `form:"file" validate:"required"`
}

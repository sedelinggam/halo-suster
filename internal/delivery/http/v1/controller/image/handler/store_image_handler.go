package imageHandler

import (
	"errors"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
	"halo-suster/package/storage/s3"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ih imageHandler) StoreImage(c echo.Context) error {

	// TODO: echo v4 not support binding multipart form data
	// Note: there is workaround for this issue, refer to https://stackoverflow.com/questions/61916842/echo-web-framework-binding-formfile

	file, err := c.FormFile("file")

	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	// Get the extension of the original file
	ext := filepath.Ext(file.Filename)

	// Generate a new filename with the UUID and the original file's extension
	filename := uuid.New().String() + ext

	// Validate the file extension
	if ext != ".jpg" && ext != ".jpeg" {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, errors.New("invalid file extension"))).SendResponse(c)
	}

	// Validate the file size
	if file.Size > 2*1024*1024 || file.Size < 10*1024 {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, errors.New("invalid file size"))).SendResponse(c)
	}

	url, err := s3.UploadFile(filename, file)

	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "File uploaded sucessfully",
		Data: response.ImageResponse{
			ImageURL: url,
		},
	})
}

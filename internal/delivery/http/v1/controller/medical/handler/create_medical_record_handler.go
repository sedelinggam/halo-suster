package medicalHandler

import (
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (mh medicalHandler) CreateMedicalRecord(c echo.Context) error {
	var (
		req  request.CreateMedicalRecord
		resp *response.CreateMedicalRecord
		err  error
	)

	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	err = mh.val.Struct(req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	resp, err = mh.medicalService.CreateMedicalRecord(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusCreated, response.Common{
		Message: "medical record created successfully",
		Data:    resp,
	})
}

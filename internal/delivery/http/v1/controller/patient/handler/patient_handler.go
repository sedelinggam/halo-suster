package patientHandler

import (
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ph patientHandler) CreatePatient(c echo.Context) error {

	var (
		req  request.CreatePatient
		resp *response.CreatePatient
		err  error
	)
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	// Validate the User struct
	err = ph.val.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	resp, err = ph.patientService.CreatePatient(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusCreated, response.Common{
		Message: "Patient created successfully",
		Data:    resp,
	})
}

package patientHandler

import (
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ph patientHandler) GetPatient(c echo.Context) error {

	var (
		req  request.PatientParam
		resp []*response.Patient
		err  error
	)
	queries := c.QueryParams()
	//Filter

	if identityNumber := queries.Get("identityNumber"); identityNumber != "" {
		identityNumber := queries.Get("identityNumber")
		req.IdentityNumber = &identityNumber
	}

	if name := queries.Get("name"); name != "" {
		name := queries.Get("name")
		req.Name = &name
	}

	if phoneNumber := queries.Get("phoneNumber"); phoneNumber != "" {
		phoneNumber := queries.Get("phoneNumber")
		phoneNumber = "+" + phoneNumber
		req.PhoneNumber = &phoneNumber
	}

	if createdAt := queries.Get("createdAt"); createdAt != "" {
		err := ph.val.Var(queries.Get("createdAt"), "oneof=asc desc")
		if err == nil {
			req.CreatedAt = &createdAt
		}
	}

	if limit := queries.Get("limit"); limit != "" {
		err := ph.val.Var(queries.Get("limit"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(queries.Get("limit"), 10, 32)
			req.Limit = int(val)
		}
	} else {
		req.Limit = 5
	}

	if offset := queries.Get("offset"); offset != "" {
		err := ph.val.Var(queries.Get("offset"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(queries.Get("offset"), 10, 32)
			req.Offset = int(val)
		}
	} else {
		req.Offset = 0
	}
	resp, err = ph.patientService.GetPatients(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	//Get jwt user ID
	return c.JSON(http.StatusOK, response.Common{
		Message: "User updated successfully",
		Data:    resp,
	})
}

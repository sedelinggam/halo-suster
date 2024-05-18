package patientHandler

import (
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ph patientHandler) GetPatient(c echo.Context) error {

	var (
		req  request.UserParam
		resp []*response.UserNurse
	)
	queries := c.QueryParams()
	//Filter

	if userId := queries.Get("userId"); userId != "" {
		userId := queries.Get("userId")
		req.UserID = &userId
	}

	if name := queries.Get("name"); name != "" {
		name := queries.Get("name")
		req.Name = &name
	}

	if nip := queries.Get("nip"); nip != "" {
		nip := queries.Get("nip")
		req.Nip = &nip
	}

	if role := queries.Get("role"); role != "" {
		err := ph.val.Var(queries.Get("role"), "oneof=it nurse")
		if err == nil {
			req.Role = &role
		}
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
	//Get jwt user ID
	return c.JSON(http.StatusOK, response.Common{
		Message: "User updated successfully",
		Data:    resp,
	})
}

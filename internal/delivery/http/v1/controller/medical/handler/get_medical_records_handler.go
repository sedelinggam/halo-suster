package medicalHandler

import (
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (mh medicalHandler) GetMedicalRecords(c echo.Context) error {

	var (
		req  request.GetMedicalRecords
		resp []*response.GetMedicalRecords
		err  error
	)
	queries := c.QueryParams()

	if idNumber := queries.Get("identityDetail.identityNumber"); idNumber != "" {
		idNumber := queries.Get("identityDetail.identityNumber")
		req.IdentityNumber = &idNumber
	}

	if userId := queries.Get("createdBy.userId"); userId != "" {
		userId := queries.Get("createdBy.userId")
		req.UserID = &userId
	}

	if nip := queries.Get("createdBy.nip"); nip != "" {
		nip := queries.Get("createdBy.nip")
		req.NIP = &nip
	}

	if limit := queries.Get("limit"); limit != "" {
		err := mh.val.Var(queries.Get("limit"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(queries.Get("limit"), 10, 32)
			req.Limit = int(val)
		}
	} else {
		req.Limit = 5
	}

	if offset := queries.Get("offset"); offset != "" {
		err := mh.val.Var(queries.Get("offset"), "number")
		if err == nil {
			val, _ := strconv.ParseInt(queries.Get("offset"), 10, 32)
			req.Offset = int(val)
		}
	} else {
		req.Offset = 0
	}

	resp, err = mh.medicalService.GetMedicalRecords(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "User updated successfully",
		Data:    resp,
	})
}

package userHandler

import (
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uh userHandler) NurseLogin(c echo.Context) error {
	var (
		req  request.UserLogin
		resp *response.UserAccessToken
		err  error
	)
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	// Validate the User struct
	err = uh.val.Struct(req)

	if err != nil {
		// Validation failed, handle the error
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)
	}

	//Role type IT
	req.RoleType = "nurse"
	resp, err = uh.userService.Login(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "User logged successfully",
		Data:    resp,
	})
}

package userHandler

import (
	"fmt"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uh userHandler) ITRegister(c echo.Context) error {
	var (
		req  request.ITUserRegister
		resp *response.UserAccessToken
		err  error
	)
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	// Validate the User struct
	fmt.Println(req.NIP, ":Z")
	err = uh.val.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	resp, err = uh.userService.RegisterUserIT(c.Request().Context(), req)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusCreated, response.Common{
		Message: "User registered successfully",
		Data:    resp,
	})
}

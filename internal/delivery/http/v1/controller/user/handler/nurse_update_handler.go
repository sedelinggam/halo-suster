package userHandler

import (
	"errors"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	cryptoJWT "halo-suster/package/crypto/jwt"
	"halo-suster/package/lumen"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (uh userHandler) NurseUpdate(c echo.Context) error {
	var (
		req    request.NurseUserUpdate
		resp   *response.UserNurse
		userId string
		err    error
	)
	if id := c.Param("userId"); id != "" {
		userId = id
	}
	err = c.Bind(&req)
	if err != nil {
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	//Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	if claims.RoleType != "it" {
		return lumen.FromError(lumen.NewError(lumen.ErrUnauthorized, errors.New("wrong user role"))).SendResponse(c)
	}

	// Validate the User struct
	err = uh.val.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		return lumen.FromError(lumen.NewError(lumen.ErrBadRequest, err)).SendResponse(c)

	}

	resp, err = uh.userService.UpdateUserNurse(c.Request().Context(), req.NIP, req.Name, userId)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Message: "User updated successfully",
		Data:    resp,
	})
}

package userHandler

import (
	"errors"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	cryptoJWT "halo-suster/package/crypto/jwt"
	"halo-suster/package/lumen"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (uh userHandler) NursePassword(c echo.Context) error {
	var (
		req  request.NurseUserPassword
		resp *response.UserNurse
		nip  int
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

	if id := c.Param("nurseId"); id != "" {
		err := uh.val.Var(c.Param("nurseId"), "int")
		if err != nil {
			return lumen.FromError(lumen.NewError(lumen.ErrNotFound, err)).SendResponse(c)
		}
		nip, _ = strconv.Atoi(id)
	}

	//Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	if claims.RoleType != "it" {
		return lumen.FromError(lumen.NewError(lumen.ErrUnauthorized, errors.New("wrong user role"))).SendResponse(c)
	}

	resp, err = uh.userService.AccessUserNurse(c.Request().Context(), nip, req.Password)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Data: resp,
	})
}

package userHandler

import (
	"errors"
	"halo-suster/internal/delivery/http/v1/response"
	cryptoJWT "halo-suster/package/crypto/jwt"
	"halo-suster/package/lumen"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (uh userHandler) NurseDelete(c echo.Context) error {
	var (
		resp    *response.UserNurse
		nurseId string
		err     error
	)

	if id := c.Param("userId"); id != "" {
		nurseId = id
	}

	//Get jwt user ID
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*cryptoJWT.JWTClaims)
	if claims.RoleType != "it" {
		return lumen.FromError(lumen.NewError(lumen.ErrUnauthorized, errors.New("wrong user role"))).SendResponse(c)
	}

	resp, err = uh.userService.DeleteUserNurse(c.Request().Context(), nurseId)
	if err != nil {
		return lumen.FromError(err).SendResponse(c)
	}

	return c.JSON(http.StatusOK, response.Common{
		Data: resp,
	})
}

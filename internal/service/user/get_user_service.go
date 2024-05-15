package userService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
)

func (ps userService) GetUsers(ctx context.Context, req request.UserParam) ([]*response.UserNurse, error) {
	products, err := ps.userRepo.GetUsers(ctx, req)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	productsResp := response.MapUserListEntityToListResponse(products)
	return productsResp, nil
}

package userRepository

import (
	"context"

	"fmt"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/entity"
	"strings"
)

func (pr userRepository) GetUsers(ctx context.Context, req request.UserParam) ([]*entity.User, error) {
	var (
		conditions []string
		filter     []interface{}
		resp       []*entity.User
		err        error
	)

	query := `SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at FROM products`

	if req.Name != nil {
		filter = append(filter, req.Name)
		conditions = append(conditions, fmt.Sprintf("name ILIKE '%%' || $%d || '%%'", len(filter)))
	}

	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
		query += " AND is_available = true AND deleted_at IS NULL"
	} else {
		query += " WHERE is_available = true AND deleted_at IS NULL"
	}

	filter = append(filter, req.Limit)
	query += fmt.Sprintf(" LIMIT $%d", len(filter))

	filter = append(filter, req.Offset)
	query += fmt.Sprintf(" OFFSET $%d", len(filter))

	err = pr.db.Select(&resp, query, filter...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

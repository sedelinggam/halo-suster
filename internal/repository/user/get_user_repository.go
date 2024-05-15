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

	query := `SELECT id, nip, name, created_at FROM users`

	if req.Name != nil {
		filter = append(filter, req.Name)
		conditions = append(conditions, fmt.Sprintf("name ILIKE '%%' || $%d || '%%'", len(filter)))
	}

	if req.Name != nil {
		filter = append(filter, req.Name)
		conditions = append(conditions, fmt.Sprintf("name ILIKE '%%' || $%d || '%%'", len(filter)))
	}

	if req.UserID != nil {
		filter = append(filter, req.UserID)
		conditions = append(conditions, fmt.Sprintf("id = $%d", len(filter)))
	}

	if req.Role != nil {
		filter = append(filter, req.Role)
		conditions = append(conditions, fmt.Sprintf("role = $%d", len(filter)))
	}

	if req.Nip != nil {
		filter = append(filter, req.Nip)
		conditions = append(conditions, fmt.Sprintf("nip ILIKE $%d || '%%'", len(filter)))
	}

	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
		query += " AND deleted_at IS NULL"
	} else {
		query += " WHERE deleted_at IS NULL"
	}

	if req.CreatedAt != nil {
		query += fmt.Sprintf(" ORDER BY created_at %s", *req.CreatedAt)
	} else {
		query += " ORDER BY created_at DESC"
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

package patientRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/entity"
	"strings"
)

func (pr patientRepository) GetPatients(ctx context.Context, req request.PatientParam) ([]*entity.Patient, error) {
	var (
		conditions []string
		filter     []interface{}
		resp       []*entity.Patient
		err        error
	)

	query := `SELECT * FROM patients`

	if req.Name != nil {
		filter = append(filter, req.Name)
		conditions = append(conditions, fmt.Sprintf("name ILIKE '%%' || $%d || '%%'", len(filter)))
	}

	if req.IdentityNumber != nil {
		filter = append(filter, req.IdentityNumber)
		conditions = append(conditions, fmt.Sprintf("identityNumber = $%d", len(filter)))
	}

	if req.PhoneNumber != nil {
		filter = append(filter, req.PhoneNumber)
		conditions = append(conditions, fmt.Sprintf("phone_number ILIKE $%d || '%%'", len(filter)))
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

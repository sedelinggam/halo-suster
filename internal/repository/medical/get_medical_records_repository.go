package medicalRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/entity"
	"strings"
)

func (pr medicalRepository) GetMedicalRecords(ctx context.Context, req request.GetMedicalRecords) ([]*entity.MedicalRecords, error) {
	var (
		conditions []string
		filter     []interface{}
		resp       []*entity.MedicalRecords
		err        error
	)

	query := `SELECT mr.*, p.phone_number, p.name as patient_name, p.birth_date, p.gender, p.identity_card_scan_url, u.nip, u.name as user_name 
			  FROM medical_records mr
			  JOIN patients p ON mr.identity_number = p.identity_number
			  JOIN users u ON mr.user_id = u.id`

	if req.IdentityNumber != nil {
		filter = append(filter, req.IdentityNumber)
		conditions = append(conditions, fmt.Sprintf("mr.identity_number = $%d", len(filter)))
	}

	if req.UserID != nil {
		filter = append(filter, req.UserID)
		conditions = append(conditions, fmt.Sprintf("mr.user_id = $%d", len(filter)))
	}

	if req.NIP != nil {
		filter = append(filter, req.NIP)
		conditions = append(conditions, fmt.Sprintf("u.nip = $%d", len(filter)))
	}

	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
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

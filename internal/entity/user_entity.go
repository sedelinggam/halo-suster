package entity

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        string      `db:"id"`
	NIP       string      `db:"nip"`
	Name      string      `db:"name"`
	Password  string      `db:"password"`
	UserRole  string      `db:"user_role"`
	CreatedAt time.Time   `db:"created_at"`
	DeletedAt pq.NullTime `db:"deleted_at"`
}

func (s User) TableName() string {
	return `users`
}

func (s User) CheckNIP() bool {
	return false
}

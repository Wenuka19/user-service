package domain

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"` // Auth0 `sub`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

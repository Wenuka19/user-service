package domain

package domain

type UserRepository interface {
	GetByID(id string) (*User, error)
	Save(user *User) error
}


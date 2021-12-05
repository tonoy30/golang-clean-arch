package repository

import "github.com/tonoy30/clean-arch/internal/domain"

type UserRepository interface {
	CreateUser(user domain.User) (*int64, error)
}

type userRepository struct {
}

func (u userRepository) CreateUser(user domain.User) (*int64, error) {
	query := pgQb().
		Insert(domain.UserTableName).
		Columns("first_name", "email", "password", "last_name",
			"role", "verified", "email_code", "balance", "phone_number").
		Values(user.FirstName, user.Email, user.Password, user.LastName,
			user.Role, user.Verified, user.EmailCode, user.Balance, user.PhoneNumber).
		Suffix("RETURNING id")

	var id int64
	err := query.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

package repository

import (
	"database/sql"
	"fmt"

	"github.com/Gergenus/internal/domain"
	"github.com/Gergenus/pkg"
)

type UserRepository interface {
	CreateUser(userName string, hashpassword string) (int, error)
	DeleteUser(userId int) (int, error)
	GetUser(name string) (*domain.User, error)
}

type postgresUserRepository struct {
	DB pkg.PostgresDatabase
}

func NewPostgresUserRepository(db pkg.PostgresDatabase) *postgresUserRepository {
	return &postgresUserRepository{DB: db}
}

func (p *postgresUserRepository) CreateUser(userName string, hashpassword string) (int, error) {
	var id int
	res := p.DB.GetDB().QueryRow("INSERT INTO users (username, hashpassword) VALUES($1, $2) RETURNING uid", userName, hashpassword)

	err := res.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *postgresUserRepository) DeleteUser(userId int) (int, error) {
	var id int
	row := p.DB.GetDB().QueryRow("DELETE FROM users WHERE uid=$1 RETURNING uid", userId)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *postgresUserRepository) GetUser(name string) (*domain.User, error) {
	row := p.DB.GetDB().QueryRow("SELECT * FROM users WHERE username=$1", name)

	var user domain.User

	err := row.Scan(&user.Uid, &user.UserName, &user.HashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/radityarestan/mini-project/internal/entity"
)

type (
	User interface {
		FetchUserById(id int64) (*entity.User, error)
		FetchUserByUsername(username string) (*entity.User, error)
		FetchUsers() ([]*entity.User, error)
		InsertUser(name string, username string, password string, role string) (*entity.User, error)
	}

	UserRepository struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserById(id int64) (*entity.User, error) {
	sqlStatement := `
		SELECT id, name, username, password, role 
		FROM USER 
		WHERE id = ?;
	`
	var user entity.User
	row := u.db.QueryRow(sqlStatement, id)

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Role,
	); err != nil {
		return &entity.User{}, err
	}

	return &user, nil
}

func (u *UserRepository) FetchUserByUsername(username string) (*entity.User, error) {
	sqlStatement := `
		SELECT id, name, username, password, role 
		FROM USER 
		WHERE username = ?;
	`
	var user entity.User
	row := u.db.QueryRow(sqlStatement, username)

	if err := row.Err(); err != nil {
		return nil, err
	}

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Role,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) FetchUsers() ([]*entity.User, error) {
	sqlStatement := `
	SELECT id, name, username, password, role 
	FROM USERS;
	`

	var users []*entity.User
	rows, err := u.db.Query(sqlStatement)
	if err != nil {
		return make([]*entity.User, 0), err
	}

	defer rows.Close()

	for rows.Next() {
		var user *entity.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.Role,
		); err != nil {
			return make([]*entity.User, 0), err
		}
	}

	return users, nil
}

func (u *UserRepository) InsertUser(name string, username string, password string, role string) (*entity.User, error) {
	sqlStatement := `
		INSERT INTO USER (name, username, password, role) 
		VALUES (?, ?, ?, ?);
	`
	result, err := u.db.Exec(sqlStatement, name, username, password, role)

	if err != nil {
		return &entity.User{}, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       userID,
		Name:     name,
		Username: username,
		Password: password,
		Role:     role,
	}, nil
}

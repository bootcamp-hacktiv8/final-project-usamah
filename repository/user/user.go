package user

import (
	"context"
	"database/sql"
	_entities "final-project-usamah/entities"
	"time"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) Register(ctx context.Context, user _entities.User) (_entities.User, int, error) {
	query := "INSERT INTO users (username, email, password, age, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := ur.database.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.Age, time.Now()).Scan(&id)
	if err != nil {
		return user, id, err
	}
	return user, id, nil
}

func (ur *UserRepository) Login(ctx context.Context, email string) (_entities.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	var user _entities.User

	err := ur.database.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return _entities.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetUser(ctx context.Context, idToken int) (_entities.User, error) {
	query := `SELECT id, username, email, age, created_at FROM users WHERE id = $1`
	var user _entities.User

	err := ur.database.QueryRowContext(ctx, query, idToken).Scan(&user.Id, &user.Username, &user.Email, &user.Age, &user.Created_at)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, updateUser _entities.User, idToken int) (_entities.User, error) {
	query := `UPDATE users SET username = $1, email = $2, age = $3, updated_at = $4
	WHERE id = $5`

	_, err := ur.database.ExecContext(ctx, query, updateUser.Username, updateUser.Email, updateUser.Age, time.Now(), idToken)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, idToken int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := ur.database.ExecContext(ctx, query, idToken)
	if err != nil {
		return err
	}
	return nil
}

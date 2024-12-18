package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/miladvatankhah/maker-checker/internal/message_approval/domain/aggregates"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) Save(user *aggregates.User) error {
	_, err := repo.db.Exec(`
        INSERT INTO users (id)
        VALUES ($1)
    `, user.ID)
	return err
}

func (repo *UserRepositoryImpl) FindByID(id uuid.UUID) (*aggregates.User, error) {
	var user aggregates.User
	err := repo.db.QueryRow(`
        SELECT id
        FROM users
        WHERE id = $1
    `, id).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

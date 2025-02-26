package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AyKrimino/JobSeekerAPI/types"
)

type userStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) types.UserRepository {
	return &userStore{
		db: db,
	}
}

func (s *userStore) GetUserByEmail(e string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM User WHERE email = ?", e)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *userStore) CreateUser(u *types.User) (int, error) {
	now := time.Now().UTC()

	res, err := s.db.Exec(
		"INSERT INTO User (email, password, role, isActive, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)",
		u.Email,
		u.Password,
		u.Role,
		true,
		now,
		now,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)

	err := rows.Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}

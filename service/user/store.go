package user

import (
	"database/sql"
	"fmt"

	"github.com/AyKrimino/JobSeekerAPI/types"
	"github.com/AyKrimino/JobSeekerAPI/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM User WHERE email = ?", email)
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

func (s *Store) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(u *types.User) (int, error) {
	res, err := s.db.Exec(
		"INSERT INTO User (email, password, role, isActive, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)",
		u.Email,
		u.Password,
		u.Role,
		u.IsActive,
		u.CreatedAt,
		u.UpdatedAt,
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

func (s *Store) CreateJobSeeker(js *types.JobSeeker) error {
	skillsJSON, err := utils.EncodeStringSliceToJSON(js.Skills)
	if err != nil {
		return fmt.Errorf("Error encoding skills to JSON: %v", err)
	}

	_, err = s.db.Exec(
		"INSERT INTO JobSeeker (firstName, lastName, profileSummary, skills, experience, education, userID) VALUES (?, ?, ?, ?, ?, ?, ?)",
		js.FirstName,
		js.LastName,
		js.ProfileSummary,
		skillsJSON,
		js.Experience,
		js.Education,
		js.UserID,
	)

	return err
}

func (s *Store) CreateCompany(cpy *types.Company) error {
	_, err := s.db.Exec(
		"INSERT INTO Company (name, headquarters, website, industry, companySize, userID) VALUES (?, ?, ?, ?, ?, ?)",
		cpy.Name,
		cpy.Headquarters,
		cpy.Website,
		cpy.Industry,
		cpy.CompanySize,
		cpy.UserID,
	)

	return err
}

func (s *Store) GetUserRoleById(id int) (string, error) {
	return "", nil
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

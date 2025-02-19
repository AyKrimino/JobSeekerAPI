package company

import (
	"database/sql"

	"github.com/AyKrimino/JobSeekerAPI/types"
)

type companyStore struct {
	db *sql.DB
}

func NewCompany(db *sql.DB) types.CompanyRepository {
	return &companyStore{
		db: db,
	}
}

func (s *companyStore) CreateCompany(cpy *types.Company) error {
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

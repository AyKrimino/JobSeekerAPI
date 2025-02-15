package jobseeker

import (
	"database/sql"
	"fmt"

	"github.com/AyKrimino/JobSeekerAPI/types"
	"github.com/AyKrimino/JobSeekerAPI/utils"
)

type jobseekerStore struct {
	db *sql.DB
}

func NewJobseekerStore(db *sql.DB) types.JobSeekerRepository {
	return &jobseekerStore{
		db: db,
	}
}

func (s *jobseekerStore) CreateJobSeeker(js *types.JobSeeker) error {
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
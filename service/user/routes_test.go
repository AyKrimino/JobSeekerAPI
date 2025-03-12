package user

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AyKrimino/JobSeekerAPI/service/company"
	"github.com/AyKrimino/JobSeekerAPI/service/jobseeker"
	"github.com/AyKrimino/JobSeekerAPI/testutils"
)

func SetupHandlerWithDB(t *testing.T) (*Handler, *sql.DB) {
	db := testutils.SetupTestDB(t)

	handler := &Handler{
		UserRepo:      NewUserStore(db),
		JobSeekerRepo: jobseeker.NewJobseekerStore(db),
		CompanyRepo:   company.NewCompany(db),
	}

	return handler, db
}

func TestHandleRegister(t *testing.T) {
	testEmail := "jobseeker@example.com"

	t.Run("Valid JobSeeker registration", func(t *testing.T) {
		handler, db := SetupHandlerWithDB(t)
		defer db.Close()

		reqBody := `{
			"email": "` + testEmail + `",
			"password": "password123",
			"role": "JobSeeker",
			"firstName": "John",
			"lastName": "Doe",
			"profileSummary": "Software Engineer",
			"skills": ["Go", "Python"],
			"experience": 2,
			"education": "BSc Computer Science"
		}`

		req := httptest.NewRequest("POST", "/register", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.handleRegister(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("expected status code 201, got %d", res.StatusCode)
		}

		createdUser, err := handler.UserRepo.GetUserByEmail(testEmail)
		if err != nil {
			t.Fatalf("failed to fetch user from DB: %v", err)
		}

		if createdUser.Email != testEmail {
			t.Fatalf("expected user email to be `%s`, got `%s`", testEmail, createdUser.Email)
		}

		if createdUser.Role != "JobSeeker" {
			t.Fatalf("expected user role to be `JobSeeker`, got `%s`", createdUser.Role)
		}
	})
}

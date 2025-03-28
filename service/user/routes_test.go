package user

import (
	"database/sql"
	"io"
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

	t.Run("Valid Company Registration", func(t *testing.T) {
		handler, db := SetupHandlerWithDB(t)
		defer db.Close()

		reqBody := `{
		"email": "` + testEmail + `",
		"password": "password123",
		"role": "Company",
		"name": "TechNova",
		"headquarters": "San Francisco, CA, USA",
		"website": "www.technova.com",
		"industry": "Information Technology",
		"companySize": "201-500"
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

		if createdUser.Role != "Company" {
			t.Fatalf("expected user role to be `Company`, got `%s`", createdUser.Role)
		}
	})

	t.Run("Invalid request body missing required role field", func(t *testing.T) {
		handler, db := SetupHandlerWithDB(t)
		defer db.Close()

		reqBody := `{
		"email": "` + testEmail + `",
		"password": "password123",
		"name": "TechNova",
		"headquarters": "San Francisco, CA, USA",
		"website": "www.technova.com",
		"industry": "Information Technology",
		"companySize": "201-500"
		}`

		req := httptest.NewRequest("POST", "/register", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.handleRegister(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status code 400, got %d", res.StatusCode)
		}

		_, err := handler.UserRepo.GetUserByEmail(testEmail)
		if err == nil {
			t.Fatal("expected user fetching to fail.")
		}
	})

	t.Run("User already exist", func(t *testing.T) {
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

		req2 := httptest.NewRequest("POST", "/register", strings.NewReader(reqBody))
		req2.Header.Set("Content-Type", "application/json")
		rec2 := httptest.NewRecorder()

		handler.handleRegister(rec2, req2)

		res2 := rec2.Result()
		defer res2.Body.Close()

		if res2.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status code 400, got %d", res2.StatusCode)
		}
	})

	t.Run("Invalid role", func(t *testing.T) {
		handler, db := SetupHandlerWithDB(t)
		defer db.Close()

		reqBody := `{
		"email": "` + testEmail + `",
		"password": "password123",
		"role": "invalid role",
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

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status code 400, got %d", res.StatusCode)
		}

		_, err := handler.UserRepo.GetUserByEmail(testEmail)
		if err == nil {
			t.Fatal("expected user fetching to fail.")
		}
	})
}

func TestHandleLogin(t *testing.T) {
	testEmail := "test@example.com"

	t.Run("Valid credentials", func(t *testing.T) {
		handler, db := SetupHandlerWithDB(t)
		defer db.Close()

		registerReqBody := `{
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

		req := httptest.NewRequest("POST", "/register", strings.NewReader(registerReqBody))
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

		loginReqBody := `{
		"email": "` + testEmail + `",
		"password": "password123"
		}`

		req2 := httptest.NewRequest("POST", "/login", strings.NewReader(loginReqBody))
		req2.Header.Set("Content-Type", "application/json")
		rec2 := httptest.NewRecorder()

		handler.handleLogin(rec2, req2)

		res2 := rec2.Result()
		defer res2.Body.Close()

		if res2.StatusCode != http.StatusOK {
			t.Fatalf("expected status code 200, got %d", res2.StatusCode)
		}

		body, _ := io.ReadAll(res2.Body)
		if !strings.Contains(string(body), "token") {
			t.Fatal("expected token to be returned")
		}
	})

	t.Run("Invalid request body missing required password field", func(t *testing.T) {
		handler, db := SetupHandlerWithDB(t)
		defer db.Close()

		loginReqBody := `{
		"email": "` + testEmail + `"
		}`

		req := httptest.NewRequest("POST", "/login", strings.NewReader(loginReqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.handleLogin(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status code 400, got %d", res.StatusCode)
		}
	})

	t.Run("Invalid password", func(t *testing.T) {
		handler, db := SetupHandlerWithDB(t)
		defer db.Close()

		registerReqBody := `{
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

		req := httptest.NewRequest("POST", "/register", strings.NewReader(registerReqBody))
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

		loginReqBody := `{
		"email": "` + testEmail + `",
		"password": "invalidPassword123"
		}`

		req2 := httptest.NewRequest("POST", "/login", strings.NewReader(loginReqBody))
		req2.Header.Set("Content-Type", "application/json")
		rec2 := httptest.NewRecorder()

		handler.handleLogin(rec2, req2)

		res2 := rec2.Result()
		defer res2.Body.Close()

		if res2.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status code 400, got %d", res2.StatusCode)
		}
	})
}

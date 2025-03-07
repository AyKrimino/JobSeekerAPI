package jobseeker_test

import (
	"testing"

	"github.com/AyKrimino/JobSeekerAPI/service/jobseeker"
	"github.com/AyKrimino/JobSeekerAPI/service/user"
	"github.com/AyKrimino/JobSeekerAPI/testutils"
	"github.com/AyKrimino/JobSeekerAPI/types"
)

func TestCreateJobSeeker_Success(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	userStore := user.NewUserStore(db)

	u := &types.User{
		Email:    "fname@test.com",
		Password: "Pass1234",
		Role:     "JobSeeker",
	}

	userID, err := userStore.CreateUser(u)
	if err != nil {
		t.Fatal("expected a nil error but got a non-nil error: ", err)
	}

	jsStore := jobseeker.NewJobseekerStore(db)

	js := &types.JobSeeker{
		FirstName:      "fname",
		LastName:       "lname",
		ProfileSummary: "psummary",
		Experience:     0,
		Education:      "edu",
		UserID:         userID,
	}

	err = jsStore.CreateJobSeeker(js)
	if err != nil {
		t.Error("expected nil error but got a non-nil error: ", err)
	}
}

func TestCreateJobSeeker_NonExistantUser(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	jsStore := jobseeker.NewJobseekerStore(db)

	js := &types.JobSeeker{
		FirstName:      "fname",
		LastName:       "lname",
		ProfileSummary: "psummary",
		Experience:     0,
		Education:      "edu",
		UserID:         0,
	}

	err := jsStore.CreateJobSeeker(js)
	if err == nil {
		t.Error("expected a nil error")
	}
}

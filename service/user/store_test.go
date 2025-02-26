package user

import (
	"testing"

	"github.com/AyKrimino/JobSeekerAPI/testutils"
	"github.com/AyKrimino/JobSeekerAPI/types"
)

func TestCreateUser_Success(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	userStore := NewUserStore(db)

	u := &types.User{
		Email:    "test@example.com",
		Password: "validpass123",
		Role:     "JobSeeker",
	}

	userID, err := userStore.CreateUser(u)
	if err != nil {
		t.Fatal("CreateUser failed:", err)
	}

	if userID == 0 {
		t.Error("expected non-zero user ID")
	}

	createdUser, err := userStore.GetUserByEmail("test@example.com")
	if err != nil {
		t.Fatal("GetUserByEmail failed:", err)
	}

	if createdUser.Email != u.Email {
		t.Errorf("expected email %s, got %s", u.Email, createdUser.Email)
	}
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	userStore := NewUserStore(db)

	user1 := &types.User{
		Email:    "duplicate@test.com",
		Password: "pass1111",
		Role:     "Company",
	}

	_, err := userStore.CreateUser(user1)
	if err != nil {
		t.Fatal("First create should succeed:", err)
	}

	user2 := &types.User{
		Email:    "duplicate@test.com",
		Password: "pass2222",
		Role:     "JobSeeker",
	}

	_, err = userStore.CreateUser(user2)
	if err == nil {
		t.Errorf("Expected error for duplicate email")
	}
}

func TestCreateUser_InvalidRole(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	userStore := NewUserStore(db)

	u := &types.User{
		Email: "invalidRole@test.com",
		Password: "pass1234",
		Role: "InvalidRole",
	}

	userID, err := userStore.CreateUser(u)
	if err == nil {
		t.Error("expected non-nil error for invalid role but got nil")
	}
	if userID != 0 {
		t.Error("expected a zero user ID when role is invalid")
	}
}

func TestCreateUser_LongEmail(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	userStore := NewUserStore(db)

	longEmail := "a" + string(make([]byte, 300)) + "@test.com"

	u := &types.User{
		Email:    longEmail,
		Password: "securepass",
		Role:     "Company",
	}

	userID, err := userStore.CreateUser(u)
	if err == nil {
		t.Error("expected error for long email but got nil")
	}
	if userID != 0 {
		t.Error("expected a zero user ID for long email")
	}
}

func TestGetUserByEmail_NotFound(t * testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	userStore := NewUserStore(db)

	_, err := userStore.GetUserByEmail("nonexistant@test.com")
	if err == nil {
		t.Error("expected error for non-existent user but got nil")
	}
}

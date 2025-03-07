package company_test

import (
	"testing"

	"github.com/AyKrimino/JobSeekerAPI/service/company"
	"github.com/AyKrimino/JobSeekerAPI/service/user"
	"github.com/AyKrimino/JobSeekerAPI/testutils"
	"github.com/AyKrimino/JobSeekerAPI/types"
)

func TestCreateCompany_Success(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	userStore := user.NewUserStore(db)

	u := &types.User{
		Email:    "test@example.com",
		Password: "validpass123",
		Role:     "Company",
	}

	userID, err := userStore.CreateUser(u)
	if err != nil {
		t.Fatal("CreateUser failed:", err)
	}

	if userID == 0 {
		t.Error("expected non-zero user ID")
	}

	cpyStore := company.NewCompany(db)

	cpy := &types.Company{
		Name:         "CompanyName",
		Headquarters: "company headquarters",
		Website:      "company.com",
		Industry:     "company industry",
		CompanySize:  "Big",
		UserID:       userID,
	}

	err = cpyStore.CreateCompany(cpy)
	if err != nil {
		t.Error("expected a nil error but got ", err)
	}
}

func TestCreateCompany_NonExistantUser(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer db.Close()

	cpyStore := company.NewCompany(db)

	cpy := &types.Company{
		Name:         "CompanyName",
		Headquarters: "company headquarters",
		Website:      "company.com",
		Industry:     "company industry",
		CompanySize:  "Big",
		UserID:       1,
	}

	err := cpyStore.CreateCompany(cpy)
	if err == nil {
		t.Error("expected a non-nil error but got a nil error")
	}
}

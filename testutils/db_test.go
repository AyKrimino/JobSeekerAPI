package testutils

import "testing"

func TestSetupTestDB(t *testing.T) {
	db := SetupTestDB(t)
	defer db.Close()
}

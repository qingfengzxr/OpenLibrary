package sqlite

import (
	"path/filepath"
	"testing"
)

var (
	testPath = filepath.Join("testdata", "db")
	testDb   = "p2pstore-test.db"
)

func TestCreateDatabaseMgr(t *testing.T) {
	dbm, err := NewDatabaseMgr(testPath, testDb)
	if err != nil || dbm == nil {
		t.Fail()
	}

	defer dbm.ORMHandler.Close()

	dbm2 := GetInstance()

	if dbm2 == nil {
		t.Fail()
	}
}

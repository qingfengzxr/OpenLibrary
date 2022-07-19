/*
 * @Author: tj
 * @Date: 2022-07-19 18:53:42
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 21:20:23
 * @FilePath: \OpenLibrary\ol\internal\sqlite\dbmgr.go
 */
package sqlite

import (
	"database/sql"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Create singleton instance for DB management
//Per history reason, we provide both GORM and Raw SQL instance

const (
	//DefaultDSN default data source name
	DefaultDSN = "openlibrarystore.db"
)

var databaseMgr *DatabaseMgr
var lock sync.RWMutex

//DatabaseMgr database manager
type DatabaseMgr struct {
	lock       sync.RWMutex
	once       sync.Once
	RawHandler *sql.DB
	ORMHandler *gorm.DB
}

//GetInstance get by process
func GetInstance() (*DatabaseMgr, error) {
	if databaseMgr == nil {
		//log.d("Get instance failed.")
		//In case dbh is not initialized and cause app crashes, create a temp db file.
		//if meta.GetAppContext().IsInitialized {
		//	databaseMgr, _ = NewDatabaseMgr(meta.GetAppContext().Config.StorageDir, meta.GetAppContext().Config.DatabaseName)
		//} else {
		mgr, err := newDatabaseMgr(DefaultDSN)
		if err != nil {
			return nil, err
		}
		databaseMgr = mgr
		//}
	}

	return databaseMgr, nil
}

//NewDatabaseMgr to be initialized while node is up
func NewDatabaseMgr(workDir, dbName string) (*DatabaseMgr, error) {
	createDirIfMissing(workDir)
	dsn := path.Join(workDir, dbName)

	return newDatabaseMgr(dsn)
}

func newDatabaseMgr(dsn string) (*DatabaseMgr, error) {
	lock.Lock()
	defer lock.Unlock()

	if databaseMgr != nil {
		return databaseMgr, nil
	}

	//Open with ORM
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	databaseMgr = &DatabaseMgr{RawHandler: db.DB(), ORMHandler: db}

	return databaseMgr, nil
}

func createDirIfMissing(dirPath string) error {
	if !strings.HasSuffix(dirPath, "/") {
		dirPath = dirPath + "/"
	}

	err := os.MkdirAll(path.Dir(dirPath), 0755)
	if err != nil {
		return err
	}

	return nil
}

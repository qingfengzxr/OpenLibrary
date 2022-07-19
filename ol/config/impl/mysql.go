/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:38
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 20:03:10
 * @FilePath: \ol\config\impl\mysql.go
 */
package impl

import (
	"errors"
	"os"
	"strings"

	"github.com/jinzhu/configor"
)

// DatabaseConfig database config
type DatabaseConfig struct {
	Host    string `yaml:"Host" required:"true"`     // databese host
	Port    int    `yaml:"Port" required:"true"`     // databese port
	User    string `yaml:"User" required:"true"`     // login user
	Passwd  string `yaml:"Password" required:"true"` // login password
	DBName  string `yaml:"DBName" required:"true"`   // connect database source name
	Charset string `yaml:"Charset"`                  // use char set
	Used    bool   `yaml:"Used"`                     // 是否启用mysql
}

// LoadFromFile load config from file
func (dc *DatabaseConfig) LoadFromFile(path string) error {
	if dc == nil {
		return os.ErrInvalid
	}

	err := configor.Load(dc, path)
	if err != nil {
		return err
	}

	dc.setDefaultValue()

	if !dc.IsValid() {
		return errors.New("invalid mysql database config")
	}

	return nil
}

func (dc *DatabaseConfig) setDefaultValue() {
	if dc.DBName == "" {
		dc.DBName = "ol"
	}

	if dc.Charset == "" {
		dc.Charset = "utf8mb4"
	}
}

// IsValid check config data valid
func (dc *DatabaseConfig) IsValid() bool {
	if dc == nil {
		return false
	}

	if dc.Host == "" || dc.Port == 0 || dc.User == "" || dc.Passwd == "" || dc.DBName == "" {
		return false
	}

	return true
}

func (dc *DatabaseConfig) HasUsed() bool {
	return dc.Used
}

// Default the config default value
func (dc *DatabaseConfig) Default(asSub bool) string {
	if asSub {
		return strings.Replace(databaseDefCfg, "\n", "\n  ", -1)
	}

	return databaseDefCfg
}

var databaseDefCfg = `
# database host
# required
Host: 127.0.0.1

# database port
# required
Port: 3306

# user to login database
# required
User: root

# password to login database
# required
# Password: root
Password: root

# the database name to connect
DBName: ol

# charactor set
Charset: utf8mb4

# whether use mysql
Used: true
`

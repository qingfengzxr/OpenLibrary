/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:38
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:46:47
 * @FilePath: \ol\config\impl\httpserver.go
 */
package impl

import (
	"errors"
	"os"
	"strings"

	"github.com/jinzhu/configor"
)

// HTTPConfig http config
type HTTPConfig struct {
	HTTPHost string `yaml:"HTTPHost" required:"true"`
	HTTPPort string `yaml:"HTTPPort" required:"true"`
}

// LoadFromFile load config from file
func (dc *HTTPConfig) LoadFromFile(path string) error {
	if dc == nil {
		return os.ErrInvalid
	}

	err := configor.Load(dc, path)
	if err != nil {
		return err
	}

	if !dc.IsValid() {
		return errors.New("invalid http server config")
	}

	return nil
}

// IsValid check config data valid
func (dc *HTTPConfig) IsValid() bool {
	if dc == nil {
		return false
	}

	if dc.HTTPHost == "" || dc.HTTPPort == "" {
		return false
	}

	return true
}

// Default the config default value
func (dc *HTTPConfig) Default(asSub bool) string {
	if asSub {
		return strings.Replace(httpDefCfg, "\n", "\n  ", -1)
	}

	return httpDefCfg
}

var httpDefCfg = `
# host
# required
HTTPHost: 0.0.0.0

# port
# required
HTTPPort: 10092
`

/*
 * @Author: tj
 * @Date: 2022-07-19 08:13:38
 * @LastEditors: tj
 * @LastEditTime: 2022-07-19 09:46:40
 * @FilePath: \ol\config\impl\basic.go
 */
package impl

import (
	"errors"
	"os"
	"strings"

	"github.com/jinzhu/configor"
)

// BasicConfig basic config
type BasicConfig struct {
}

// LoadFromFile load config from file
func (dc *BasicConfig) LoadFromFile(path string) error {
	if dc == nil {
		return os.ErrInvalid
	}

	err := configor.Load(dc, path)
	if err != nil {
		return err
	}

	dc.setDefaultValue()

	if !dc.IsValid() {
		return errors.New("invalid BasicConfig config")
	}

	return nil
}

func (dc *BasicConfig) setDefaultValue() {
}

// IsValid check config data valid
func (dc *BasicConfig) IsValid() bool {
	if dc == nil {
		return false
	}

	return true
}

// Default the config default value
func (dc *BasicConfig) Default(asSub bool) string {
	if asSub {
		return strings.Replace(basicDefCfg, "\n", "\n  ", -1)
	}

	return basicDefCfg
}

var basicDefCfg = ``

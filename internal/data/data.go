// Package data performs application configuration functions
package data

import (
	"os"

	"github.com/adrg/xdg"
)

// DataDir describes the path on the local system which the application can use to store data
var DataDir string = xdg.DataHome + "/oats/"

// MkDataDir checks for the existance of the application's expected data storage location
// and creates it if it does not already exist
func MkDataDir() error {
	err := os.MkdirAll(DataDir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

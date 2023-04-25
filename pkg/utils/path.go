package utils

import (
	"errors"
	"os"
)

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(path, 0o777); err != nil {
			return err
		}
	}
	return nil
}

func IsNotExistsDirectory(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return true
	}
	return false
}

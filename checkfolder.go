// checkfolder.go
package main

import (
	"os"
	"path"
)

func CheckFolder(filename string) error {
	dir, _ := path.Split(filename)
	if dir != "" {
		err := os.MkdirAll(dir, filePERM)
		return err
	}
	return nil
}
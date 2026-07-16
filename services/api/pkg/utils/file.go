package utils

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func FileExecuteDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Fatal(err)
	}
	return dir
}

package utils

import (
	"os"
	"path/filepath"
	"trekkstay/pkgs/log"
)

func GetFileName(filePath string) string {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.JsonLogger.Error(err.Error())
		panic(err)
	}

	return fileInfo.Name()
}

func GetDirectoryPath(filePath string) string {
	dirPath := filepath.Dir(filePath)
	return dirPath
}

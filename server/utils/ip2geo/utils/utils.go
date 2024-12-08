package utils

import (
	"os"
	"runtime"
	"strings"
)

func InitDataPath(replacePath, appendPath string) string {
	file := strings.Replace(getWorkPath(), replacePath, "", 1)
	return WindowsDirSpitReplace(file + appendPath)
}

func WindowsDirSpitReplace(dataPath string) string {
	if runtime.GOOS == "windows" {
		dataPath = strings.ReplaceAll(dataPath, "/", "\\")
	}
	return dataPath
}

func getWorkPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

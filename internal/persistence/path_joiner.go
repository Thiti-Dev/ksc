package persistence

import (
	"os"
	"path/filepath"
)

func JoinOriginPath(filename string) string {
	exePath, _ := os.Executable()

	exeDir := filepath.Dir(exePath)

	return filepath.Join(exeDir, "/"+filename)

}

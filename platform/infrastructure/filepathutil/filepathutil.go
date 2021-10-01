package filepathutil

import (
	"os"
	"path/filepath"
)

//JoinWithRootDir join a relative path with the root directory
func JoinWithRootDir(relativePath string) (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(rootDir, relativePath), nil
}

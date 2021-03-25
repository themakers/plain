package storage

import (
    "os"
	"path/filepath"
	"strings"
)

func shellExpand(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		path = filepath.Join(home, path[2:])
	}

	return os.ExpandEnv(path)
}

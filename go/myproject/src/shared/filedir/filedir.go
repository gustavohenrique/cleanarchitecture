package filedir

import (
	"os"
	"path/filepath"
)

func WriteFile(filename string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}

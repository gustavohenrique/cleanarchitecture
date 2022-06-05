package fileutils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Sed struct {
	repoDir    string
	distDir      string
	extensions   []string
	placeholders map[string]string
}

func NewSed() *Sed {
	return &Sed{}
}

func (pt *Sed) From(dir string) *Sed {
	pt.repoDir = dir
	return pt
}

func (pt *Sed) To(dir string) *Sed {
	pt.distDir = dir
	return pt
}

func (pt *Sed) Only(extensions []string) *Sed {
	pt.extensions = extensions
	return pt
}

func (pt *Sed) Replace(placeholders map[string]string) *Sed {
	pt.placeholders = placeholders
	return pt
}

func (pt *Sed) Run() (string, error) {
	err := filepath.Walk(pt.repoDir, pt.walkDirFn)
	return pt.distDir, err
}

func (pt *Sed) walkDirFn(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		logInfo(err)
		return err
	}
	if !fileInfo.IsDir() {
		filename := fileInfo.Name()
		if isInvalidFile(filename, pt.extensions) {
			newDir := pt.getDistDirFrom(path)
			pt.mkdir(newDir)
			return pt.copy(path, filepath.Join(newDir, filename))
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			logInfo("Failed to read repo file.", path, ". ERROR:", err)
			return err
		}
		tpl, err := template.New("").Parse(string(content))
		if err != nil {
			logInfo("Failed to parse template", path, ". ERROR:", err)
			return nil
		}
		var parsed bytes.Buffer
		err = tpl.Execute(&parsed, pt.placeholders)
		if err != nil {
			logInfo("Failed to execute template.", err)
			return err
		}
		newDir := pt.getDistDirFrom(path)
		pt.mkdir(newDir)
		newFile := filepath.Join(newDir, filename)
		return ioutil.WriteFile(newFile, parsed.Bytes(), 0644)
	}
	return nil
}

func (pt *Sed) getDistDirFrom(path string) string {
	baseDir := filepath.Dir(path)
	newDir := strings.ReplaceAll(baseDir, pt.repoDir, pt.distDir)
	return newDir
}

func (pt *Sed) mkdir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		logInfo("Failed to create dist dir:", path, ". ERROR:", err)
	}
	return err
}

func (pt *Sed) copy(src, dst string) error {
	repoFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !repoFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	repo, err := os.Open(src)
	if err != nil {
		return err
	}
	defer repo.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := repo.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

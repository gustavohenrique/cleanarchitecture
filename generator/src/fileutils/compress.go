package fileutils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const BUFFERSIZE = int64(20)

type Compress struct {
	inputDir string
	outputDir string
	outputFile  string
	excludeDirs []string
}

func NewCompress() *Compress {
	return &Compress{}
}

func (cp *Compress) Input(dir string) *Compress {
	cp.inputDir = dir
	return cp
}

func (cp *Compress) Output(dir string) *Compress {
	cp.outputDir = dir
	return cp
}

func (cp *Compress) getOutputFile() string {
	return filepath.Join(cp.outputDir, cp.outputFile)
}

func (cp *Compress) Name(name string) *Compress {
	cp.outputFile = name
	return cp
}

func (cp *Compress) Exclude(dirs []string) *Compress {
	cp.excludeDirs = dirs
	return cp
}

func (cp *Compress) Run() (string, error) {
	inputDir := cp.inputDir
	excludeDirs := cp.excludeDirs
	outputFile := cp.getOutputFile()
	tarFile, err := runTar(inputDir, outputFile, excludeDirs)
	if err != nil {
		return "", err
	}
	gzFile, err := runGzip(tarFile)
	return gzFile, err
}

func runTar(inputDir, outputFile string, excludeDirs []string) (string, error) {
	tarfilePath := fmt.Sprintf("%s.tar", outputFile)
	tarfile, err := os.Create(tarfilePath)
	if err != nil {
		logInfo("Failed to create", outputFile)
		return "", err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(inputDir)
	if err != nil {
		logInfo("Failed to get stat of source dir", inputDir, ". ERROR:", err)
		return "", nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(inputDir)
	}

	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logInfo("Failed to walk. ERROR:", err)
			return err
		}
		dirNameShouldBeIgnored := sliceContainsDir(excludeDirs, info.Name())
		pathShouldBeIgnore := contains(path, excludeDirs)
		if dirNameShouldBeIgnored || pathShouldBeIgnore {
			return nil
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			logInfo("Failed to get header info from", info.Name(), ". ERROR:", err)
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, inputDir))
		}

		if err := tarball.WriteHeader(header); err != nil {
			logInfo("Failed to write header. ERROR:", err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			logInfo("Failed to open", path, ". ERROR:", err)
			return err
		}
		defer file.Close()
		_, err = io.Copy(tarball, file)
		return err
	})
	return tarfilePath, err
}

func runUntar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func runGzip(source string) (string, error) {
	reader, err := os.Open(source)
	if err != nil {
		logInfo("Failed to open tar file", source)
		return "", err
	}
	filename := filepath.Base(source)
	parent := strings.ReplaceAll(filepath.Clean(source), filename, "")
	gzfilePath := filepath.Join(parent, fmt.Sprintf("%s.gz", filename))
	writer, err := os.Create(gzfilePath)
	if err != nil {
		logInfo("Failed to gzip", gzfilePath)
		return "", err
	}
	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	_, err = io.Copy(archiver, reader)
	return gzfilePath, err
}

func runUnGzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}

func dirSize(path string) int64 {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		logInfo("Failed to get dir size", path)
	}
	return size
}

package file

import (
	"bytes"
	"github.com/Masterminds/sprig/v3"
	"github.com/leftbin/go-util/pkg/shell"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

// IsFileExists check if a file exists
func IsFileExists(f string) bool {
	if f == "" {
		return false
	}
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsDirExists check if a directory exists
func IsDirExists(d string) bool {
	if d == "" {
		return false
	}
	info, err := os.Stat(d)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// Unzip unzips file using shell command
// expects unzip package on the os
func Unzip(zipFile, dest string) error {
	unzipCmd := exec.Command("unzip", zipFile)
	unzipCmd.Dir = dest

	if err := shell.RunCmd(unzipCmd); err != nil {
		return errors.Wrap(err, "failed to unzip")
	}
	return nil
}

func RenderTmplt(input interface{}, tmpltString, outputPath string) ([]byte, error) {
	log.Infof("rendering %s", outputPath)
	t := template.New("template").Funcs(template.FuncMap(sprig.FuncMap()))
	t, _ = t.Parse(tmpltString)
	var renderedBytes bytes.Buffer
	if err := t.Execute(&renderedBytes, input); err != nil {
		return nil, errors.Wrapf(err, "failed to render %s", outputPath)
	}
	return renderedBytes.Bytes(), nil
}

func WriteFile(content []byte, outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to create %s directory", filepath.Dir(outputPath))
	}
	err := os.WriteFile(outputPath, content, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "failed to write %s file", outputPath)
	}
	return nil
}

func Download(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

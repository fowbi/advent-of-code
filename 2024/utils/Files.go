package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Filename is the Node.js __filename equivalent
// see https://stackoverflow.com/a/70491592
func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

// Dirname is the Node.js __dirname equivalent
// see https://stackoverflow.com/a/70491592
func Dirname() (string, error) {
	filename, err := Filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}

func StarFolder(day int) string {
	cwd, cwdErr := Dirname()
	if cwdErr != nil {
		panic("Could not determine CWD")
	}

	return filepath.Join(cwd, "..", fmt.Sprintf("star_%02d", day))
}

func TemplateFolder() string {
	cwd, cwdErr := Dirname()
	if cwdErr != nil {
		panic("Could not determine CWD")
	}

	return filepath.Join(cwd, "..", "template")
}

func CreateStarFolder(day int) string {
	starFolder := StarFolder(day)

	// Make sure we do not override the star folder
	if _, err := os.Stat(starFolder); !os.IsNotExist(err) {
		log.Fatalf("Star directory already exists for day %d: %v", day, err)
	}

	err := os.MkdirAll(starFolder, os.ModePerm)
	if err != nil {
		log.Fatalf("Star directory could not be created for day %d: %v", day, err)
	}

	return starFolder
}

func WriteInput(filename string, contents []byte) error {
	err := os.WriteFile(filename, contents, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("file could not be written: %s", err)
	}

	return err
}

func CopyTemplate(day int) (string, error) {
	templateFolder := TemplateFolder()
	starFolder := CreateStarFolder(day)
	files := []string{"main", "main_test"}

	for _, file := range files {
		srcFile, err := os.Open(filepath.Join(templateFolder, file))
		if err != nil {
			return starFolder, fmt.Errorf("could not open/find %s: %w", filepath.Join(templateFolder, file), err)
		}
		defer srcFile.Close()

		if _, err := os.Stat(filepath.Join(starFolder, file+".go")); !os.IsNotExist(err) {
			return starFolder, fmt.Errorf("%s already exists: %v", file+".go", err)
		}

		destFile, err := os.Create(filepath.Join(starFolder, file+".go"))
		if err != nil {
			return starFolder, fmt.Errorf("failed to create %s: %w", filepath.Join(starFolder, file+".go"), err)
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return starFolder, fmt.Errorf("failed to copy from template to star folder: %w", err)
		}
	}

	return starFolder, nil
}

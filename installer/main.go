package main

import (
	"archive/zip"
	_ "embed"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

//go:embed files/name.txt
var ProgramName string

//go:embed files/packedFiles.zip
var ZipData []byte

// DataStruct contains the saved files and dirs
type DataStruct struct {
	SavedFiles map[string][]File
	SavedDirs  []Directory
}

// Directory represents a directory
type Directory struct {
	Path       string
	Permission os.FileMode
}

// File represents a local file
type File struct {
	Name       string
	Path       string
	Bytes      []byte
	Permission os.FileMode
}

// ZipName : the temporary zip's name
var ZipName = "packedFiles.zip"

// Program
func main() {
	RecoverFileStructure()
}

// CreateFiles creates the file structure saved in files in another location
func CreateFiles(installDir string) {
	os.MkdirAll(installDir, os.ModePerm)

	// temporarely save zip file
	zipPath := fmt.Sprint(installDir, string(filepath.Separator), ZipName)
	err := ioutil.WriteFile(zipPath, ZipData, 0744)

	if err != nil {
		panic(err)
	}

	// Unpack zip file
	err = UnpackZip(zipPath, installDir)
	if err != nil {
		panic(err)
	}

	// Delete zip file
	os.Remove(zipPath)
}

// UnpackZip unpacks the archive in the given path
func UnpackZip(archive, target string) error {

	var fileReader io.ReadCloser
	var targetFile *os.File

	// Create reader
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	// Create directories
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	// Loop and create files
	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err = file.Open()
		if err != nil {
			return err
		}

		targetFile, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	// Close
	reader.Close()
	fileReader.Close()
	targetFile.Close()

	return nil
}

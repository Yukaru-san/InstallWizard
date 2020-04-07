package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
)

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

var (
	// Files contains all files sorted by directory
	Files map[string][]File
	// Dirs contains all dirs
	Dirs []Directory

	// TempDir temporarely holds all data needed
	TempDir = "Temp %$% (do not delete!)"
)

// SaveInstallerFiles puts the main files from the explorer into the temp directory TODO Add add files needed
func SaveInstallerFiles() error {
	var err error

	// Open file compiled into the exe TODO
	box := packr.NewBox("installer")

	mainFile, err := box.Open("main.go")
	mainData, err := ioutil.ReadAll(mainFile)

	recoverFile, err := box.Open("recover.go")
	recoverData, err := ioutil.ReadAll(recoverFile)

	err = ioutil.WriteFile(fmt.Sprint(TempDir, string(filepath.Separator), "main.go"), mainData, 744)
	err = ioutil.WriteFile(fmt.Sprint(TempDir, string(filepath.Separator), "recover.go"), recoverData, 744)

	return err
}

// SaveDataAsJSON will write the json into the pack directory
func SaveDataAsJSON() error {

	// Create JSON
	data := DataStruct{SavedFiles: Files, SavedDirs: Dirs}
	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	// Save to temp directory
	err = ioutil.WriteFile(fmt.Sprint(TempDir, string(filepath.Separator), "explorerData.json"), json, 744)

	if err != nil {
		return err
	}

	return nil
}

// Explore explores a path.
// If path is empty, it explores the current directory
func Explore(path string) error {

	Files = make(map[string][]File)
	Dirs = []Directory{}

	if len(path) == 0 {
		path = "."
	}

	filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && !IsIgnored(path, info) {
				split := strings.Split(path, string(filepath.Separator))
				var filePath string
				if len(split) == 1 {
					filePath = "."
				} else {
					filePath = strings.Join(split[0:len(split)-1], string(filepath.Separator))
				}
				bytes, _ := ioutil.ReadFile(path)

				Files[filePath] = append(Files[filePath], File{Name: info.Name(), Path: filePath, Bytes: bytes, Permission: info.Mode()})
			} else if !IsIgnored(path, info) {
				Dirs = append(Dirs, Directory{Path: path, Permission: info.Mode()})
			}

			return nil
		})
	return nil
}

// IsIgnored returns true if the file or folder should be ignored when packing
func IsIgnored(path string, info os.FileInfo) bool {
	// TODO

	if strings.Contains(path, TempDir) {
		return true
	}

	return false
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	TempDir = "Temp (do not delete!)"
)

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

			if !info.IsDir() {
				split := strings.Split(path, string(filepath.Separator))
				var filePath string
				if len(split) == 1 {
					filePath = "."
				} else {
					filePath = strings.Join(split[0:len(split)-1], string(filepath.Separator))
				}
				bytes, _ := ioutil.ReadFile(path)

				Files[filePath] = append(Files[filePath], File{Name: info.Name(), Path: filePath, Bytes: bytes, Permission: info.Mode()})
			} else {
				Dirs = append(Dirs, Directory{Path: path, Permission: info.Mode()})
			}

			return nil
		})
	return nil
}

// CreateFiles creates the file structure saved in files in another location
func CreateFiles(newBasePath string) {
	if newBasePath[len(newBasePath)-1] == filepath.Separator {
		newBasePath = newBasePath[0 : len(newBasePath)-2]
	}

	for _, dir := range Dirs {
		os.MkdirAll(fmt.Sprint(newBasePath, string(filepath.Separator), dir.Path), dir.Permission)

		for _, file := range Files[dir.Path] {
			var err error
			if file.Path == "." {
				err = ioutil.WriteFile(fmt.Sprint(newBasePath, string(filepath.Separator), file.Name), file.Bytes, file.Permission)
			} else {
				err = ioutil.WriteFile(fmt.Sprint(newBasePath, string(filepath.Separator), file.Path, string(filepath.Separator), file.Name), file.Bytes, file.Permission)
			}
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

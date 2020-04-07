package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

var data DataStruct

// RecoverFileStructure recovers files into DataStruct
func RecoverFileStructure() {
	// json file
	box := packr.NewBox(".")
	jsonFile, err := box.Open("explorerData.json")
	jsonData, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
	}
}

// CreateFiles creates the file structure saved in files in another location
func CreateFiles() {

	// New Path
	box := packr.NewBox(".")
	dirFile, err := box.Open("baseDir")
	dirData, err := ioutil.ReadAll(dirFile)

	if err != nil {
		fmt.Println(err.Error())
	}

	newBasePath := string(dirData)

	if newBasePath[len(newBasePath)-1] == filepath.Separator {
		newBasePath = newBasePath[0 : len(newBasePath)-2]
	}

	for _, dir := range data.SavedDirs {
		os.MkdirAll(fmt.Sprint(newBasePath, string(filepath.Separator), dir.Path), dir.Permission)

		for _, file := range data.SavedFiles[dir.Path] {
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

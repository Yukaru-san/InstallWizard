package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gobuffalo/packr"
)

/*
	TODO
		Check for args
		Cross-Platform directory args
		Ignore desired files / folders

*/

func main() {

	if os.Args[1] == "" {
		fmt.Println("first arg should be your desired directory")
		return
	}

	fmt.Println("Creating temporary directory")
	err := os.Mkdir(TempDir, 0744)
	printError(err)

	fmt.Println("Exploring directory")
	err = Explore("")
	printError(err)

	fmt.Println("Saving data structure")
	err = SaveDataAsJSON()
	printError(err)

	fmt.Println("Saving intaller files")
	err = SaveInstallerFiles()
	printError(err)

	fmt.Println("Creating baseDir directions")
	ioutil.WriteFile(fmt.Sprint(TempDir, string(filepath.Separator), "baseDir"), []byte(os.Args[1]), 0744)

	fmt.Println("Creating packr")
	// Create Packr binary
	box := packr.NewBox("packr")
	var packrFile http.File
	var packrData []byte
	var binaryName string

	switch runtime.GOOS {
	case "windows":
		binaryName = "packr.exe"
		packrFile, err = box.Open(binaryName)
		packrData, err = ioutil.ReadAll(packrFile)
	case "darwin":
		binaryName = "packr_Mac"
		packrFile, err = box.Open(binaryName)
		packrData, err = ioutil.ReadAll(packrFile)
	case "linux":
		binaryName = "packr_Linux"
		packrFile, err = box.Open(binaryName)
		packrData, err = ioutil.ReadAll(packrFile)
	}
	printError(err)

	err = ioutil.WriteFile(fmt.Sprint(TempDir, string(filepath.Separator), binaryName), packrData, 0744)
	printError(err)

	fmt.Println("Creating new installer...")
	// Set up Paths
	// splitPath := strings.Split(os.Args[0], string(filepath.Separator))
	// filePath := strings.Join(splitPath[0:len(splitPath)-1], string(filepath.Separator))
	//	tempDir := fmt.Sprint(filePath, string(filepath.Separator), TempDir)

	// Create installer binary
	var cmd *exec.Cmd
	var installerName string

	if runtime.GOOS == "windows" {
		installerName = "installer.exe"
		cmd = exec.Command(binaryName, "build", "-o", installerName)
	} else {
		installerName = "installer"
		cmd = exec.Command(binaryName, "build", "-o", installerName)
	}

	cmd.Dir = TempDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	cmd.Wait()

	if err != nil {
		printError(err)
	}

	fmt.Println("Finalising...")
	// Create output
	os.Mkdir("output", 700)

	binary, err := ioutil.ReadFile(fmt.Sprint(TempDir, string(filepath.Separator), installerName))

	if err != nil {
		printError(err)
	}

	err = ioutil.WriteFile(fmt.Sprint("output", string(filepath.Separator), installerName), binary, 0744)

	if err != nil {
		printError(err)
	}

	// Clean up
	err = os.RemoveAll(TempDir)

	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*
Windows: GOOS=windows GOARCH=amd64
MAC:  GOOS=darwin GOARCH=amd64 go build


fmt.Println("Saved in:", fmt.Sprint(gaw.GetHome(), string(filepath.Separator), ".dmanager", string(filepath.Separator), "TestDirectory"))

	err := Explore("")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		CreateFiles(fmt.Sprint(gaw.GetHome(), string(filepath.Separator), ".dmanager", string(filepath.Separator), "TestDirectory"))
	}
*/

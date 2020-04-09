package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
	TempDir = "WorkingDirectoryDoNotDelete"
	zipName = "packedFiles.zip"

	// directions needed for save and cleanup
	zipDir                = ""
	packrAlreadyInstalled = false

	// ExecutableName is the process's name that started this program
	ExecutableName = ""

	// ZipArchive is the zip's archive
	ZipArchive *zip.Writer
)

// ImplementFiles searches the given dir and implements found files
func ImplementFiles(sourcePath string) error {
	zipfile, err := os.Create(fmt.Sprint(TempDir, string(filepath.Separator), zipName))
	if err != nil {
		return err
	}
	defer zipfile.Close()

	ZipArchive = zip.NewWriter(zipfile)
	defer ZipArchive.Close()

	info, err := os.Stat(sourcePath)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(sourcePath)
	}

	fmt.Println("---starting to searching---")
	filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {

		// Check if file should be ignored
		if IsIgnored(path, info) {
			fmt.Println("   - ignoring", info.Name())
			return nil
		}
		fmt.Println("   + implementing", info.Name())

		if err != nil {
			return err
		}

		// Set file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Handle directory entries
		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, sourcePath))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		// Create file entry and fill it
		writer, err := ZipArchive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
	fmt.Println("---finished searching---")

	return err
}

// ImplementInstallerFiles puts the main files from the explorer into the temp directory TODO Add add files needed
func ImplementInstallerFiles() error {
	var err error

	// Open file compiled into the exe (.inst, cause packr causes issues otherwise)
	box := packr.NewBox("installer")

	mainFile, err := box.Open("main.inst")
	mainData, err := ioutil.ReadAll(mainFile)

	recoverFile, err := box.Open("recover.inst")
	recoverData, err := ioutil.ReadAll(recoverFile)

	err = ioutil.WriteFile(fmt.Sprint(TempDir, string(filepath.Separator), "main.go"), mainData, 0744)
	err = ioutil.WriteFile(fmt.Sprint(TempDir, string(filepath.Separator), "recover.go"), recoverData, 0744)

	return err
}

// ImplementPackrLibrary implements the data from packr
func ImplementPackrLibrary() error {

	// Checking if packr's library is installed on the system
	githubPath := fmt.Sprint(os.Getenv("GOPATH"), string(filepath.Separator), "src", string(filepath.Separator), "github.com", string(filepath.Separator))
	buffaloDir := fmt.Sprint(os.Getenv("GOPATH"), string(filepath.Separator), "src", string(filepath.Separator), "github.com", string(filepath.Separator), "gobuffalo", string(filepath.Separator))
	packrDir := fmt.Sprint(os.Getenv("GOPATH"), string(filepath.Separator), "src", string(filepath.Separator), "github.com", string(filepath.Separator), "gobuffalo", string(filepath.Separator), "packr", string(filepath.Separator))

	var err error

	fmt.Println(" > checking for packr library inside GOPATH")
	if _, err := os.Stat(packrDir); !os.IsNotExist(err) {
		fmt.Println("   -> library found. Not installing.")
	} else {
		fmt.Println("   -> library not found")
		fmt.Println("      -> installing now")

		// Get zip file
		box := packr.NewBox("packr")
		packrZipFile, err := box.Open("packr_lib.zip")
		packrZipBytes, err := ioutil.ReadAll(packrZipFile)
		// Save it to the drive
		err = os.MkdirAll(buffaloDir, 0744)
		zipDir = fmt.Sprint(githubPath, string(filepath.Separator), "packrLib.zip")
		err = ioutil.WriteFile(zipDir, packrZipBytes, 0744)
		// Unpack
		err = unpackZip(zipDir, buffaloDir)
		if err != nil {
			return err
		}

	}

	// Check if packr's binary is already installed on the system
	fmt.Println(" > trying to find packr binary")

	findCmd := exec.Command("packr")
	err = findCmd.Run()
	findCmd.Wait()

	if err == nil {
		fmt.Println("   -> Packr is already installed.")
		return nil
	}
	packrAlreadyInstalled = true

	// Build packr binary
	fmt.Println(" > building packr binary")
	buildCmd := exec.Command("go", "build")
	buildCmd.Dir = fmt.Sprint(packrDir, string(filepath.Separator), "packr")
	err = buildCmd.Run()
	buildCmd.Wait()

	if err != nil {
		return err
	}

	// Add packr to path (enabled accessing by command)
	fmt.Println(" > adding packr binary to $PATH (temporarily)")
	var addCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		addCmd = exec.Command("packr.exe", "install")
	} else {
		addCmd = exec.Command("./packr", "install")
	}
	addCmd.Dir = fmt.Sprint(packrDir, string(filepath.Separator), "packr")
	err = addCmd.Run()
	addCmd.Wait()

	return err
}

// BuildNewBinary builds the installer binary (supports: windows, linux, darwin)
func BuildNewBinary(targetSystem string) error {
	var cmd *exec.Cmd
	var installerName string

	for i := 0; i < 3; i++ {
		if i == 0 {
			os.Setenv("GOOS", "linux")
		} else if i == 1 {
			os.Setenv("GOOS", "windows")
		} else {
			os.Setenv("GOOS", "darwin")
		}

		if os.Getenv("GOOS") == "windows" {
			installerName = "WindowsInstaller.exe"
			cmd = exec.Command("packr", "build", "-o", installerName)
		} else if os.Getenv("GOOS") == "linux" {
			installerName = "LinuxInstaller"
			cmd = exec.Command("packr", "build", "-o", installerName)
		} else {
			installerName = "DarwinInstaller"
			cmd = exec.Command("packr", "build", "-o", installerName)
		}

		cmd.Dir = TempDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		cmd.Wait()

		if err != nil {
			return err
		}

		// Create output dir
		if i == 0 {
			os.Mkdir("output", 0744)
		}

		// Read binary
		tempBinaryPath := fmt.Sprint(TempDir, string(filepath.Separator), installerName)
		binary, err := ioutil.ReadFile(tempBinaryPath)

		if err != nil {
			return err
		}

		// Save binary
		binaryPath := fmt.Sprint("output", string(filepath.Separator), installerName)
		err = ioutil.WriteFile(binaryPath, binary, 0744)

		// Delete the temporary binary
		os.Remove(tempBinaryPath)

		if err != nil {
			return err
		}

	}

	fmt.Print("\nDone! Check your output directory.\n\n")

	return nil
}

// CleanUp deletes all temp files
func CleanUp() error {

	// Temp dir
	err := os.RemoveAll(TempDir)

	// zip
	os.Remove(zipDir)

	// binary
	if packrAlreadyInstalled {
		if runtime.GOOS == "windows" {
			err = os.Remove(fmt.Sprint(os.Getenv("GOPATH"), string(filepath.Separator), "bin", string(filepath.Separator), "packr.exe"))
		} else {
			err = os.Remove(fmt.Sprint(os.Getenv("GOPATH"), string(filepath.Separator), "bin", string(filepath.Separator), "packr"))
		}
	}

	return err
}

// IsIgnored returns true if the file or folder should be ignored when packing
func IsIgnored(path string, info os.FileInfo) bool {
	// TODO
	if strings.Contains(path, TempDir) || strings.Contains(path, "ExecutableName") || strings.Contains(path, ".git") || strings.Contains(path, "output") || strings.Contains(path, "packr") {
		return true
	}

	return false
}

func unpackZip(archive, target string) error {

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

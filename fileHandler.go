package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "embed"
)

//go:embed installer/recover.go
var guiRecover []byte

//go:embed installer/recover_cli.go
var cliRecover []byte

//go:embed installer/main.go
var mainFile []byte

//go:embed installer/go.mod.x
var modFile []byte

//go:embed installer/go.sum.x
var sumFile []byte

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

	// TempDir holds all data needed
	TempDir = ".WorkingDirectoryDoNotDelete"
	// FilesPath holds installer data
	FilesPath = ""
	zipName   = "packedFiles.zip"

	// directions needed for save and cleanup
	zipDir = ""

	// ExecutableName is the process's name that started this program
	ExecutableName = ""

	// ZipArchive is the zip's archive
	ZipArchive *zip.Writer
)

// ImplementFiles searches the given dir and implements found files
func ImplementFiles(sourcePath string) error {
	// Create zip folder path
	FilesPath = fmt.Sprint(filepath.Join(TempDir, "files"))
	os.Mkdir(FilesPath, 0744)

	zipfile, err := os.Create(filepath.Join(FilesPath, zipName))
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

	fmt.Println("---starting to search---")
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

// ImplementInstallerName tells the installer it's name
func ImplementInstallerName() error {

	name := ""

	if len(os.Args) == 1 {

		fmt.Print("\nYou didn't specify a name when starting the program.\nPlease do so now:\n")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = strings.ReplaceAll(input, "\r\n", "")
		input = strings.ReplaceAll(input, "\n", "")

		name = input
	} else {
		name = os.Args[1]
	}

	fmt.Print("\n\n")
	err := os.WriteFile(fmt.Sprint(FilesPath, string(filepath.Separator), "name.txt"), []byte(name), 0744)
	return err
}

// ImplementInstallerFiles puts the main files from the explorer into the temp directory TODO Add add files needed
func ImplementInstallerFiles() error {
	var err error

	err = os.WriteFile(filepath.Join(TempDir, "recover.go"), guiRecover, 0744)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(TempDir, "recover_cli.go"), cliRecover, 0744)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(TempDir, "main.go"), mainFile, 0744)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(TempDir, "go.sum"), sumFile, 0744)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(TempDir, "go.mod"), modFile, 0744)

	return err
}

// ImplementSqweekLibrary downloads needed libs if needed
func ImplementSqweekLibrary() error {

	// Install Sqweek if needed
	var err error
	fmt.Println(" > installing sqweek if needed")
	installCmd := exec.Command("go", "get", "github.com/gen2brain/dlgs")
	err = installCmd.Run()
	installCmd.Wait()

	// Sqweek has problems using cross-platform building - installing missing libraries
	fmt.Println(" > installing w32 if needed")
	installCmd = exec.Command("go", "get", "github.com/TheTitanrain/w32")
	installCmd.Run()
	installCmd.Wait()

	return err
}

// BuildNewBinary builds the installer binary (supports: windows, linux, darwin)
func BuildNewBinary() error {
	var cmd *exec.Cmd
	var installerName string

	// Build command
	for i := 0; i < 3; i++ {
		if i == 0 {
			os.Setenv("GOOS", "windows")
		} else if i == 1 {
			os.Setenv("GOOS", "darwin")
		} else {
			/*
				// linux needs another "recover" file
				os.Remove(fmt.Sprint(TempDir, string(filepath.Separator), "main.go"))
				os.Rename(fmt.Sprint(TempDir, string(filepath.Separator), "main.linux"), fmt.Sprint(TempDir, string(filepath.Separator), "main.go"))
			*/
			os.Setenv("GOOS", "linux")
		}

		if os.Getenv("GOOS") == "windows" {
			installerName = "WindowsInstaller.exe"
			cmd = exec.Command("go", "build", "-o", installerName, "-ldflags=-s -w -H=windowsgui")
		} else if os.Getenv("GOOS") == "linux" {
			installerName = "LinuxInstaller"
			cmd = exec.Command("go", "build", "-o", installerName, "-ldflags=-s -w")
		} else {
			installerName = "DarwinInstaller"
			cmd = exec.Command("go", "build", "-o", installerName, "-ldflags=-s -w")
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
		binary, err := os.ReadFile(tempBinaryPath)

		if err != nil {
			return err
		}

		// Save binary
		binaryPath := fmt.Sprint("output", string(filepath.Separator), installerName)
		err = os.WriteFile(binaryPath, binary, 0744)

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

	return err
}

// IsIgnored returns true if the file or folder should be ignored when packing
func IsIgnored(path string, info os.FileInfo) bool {
	// TODO
	if strings.Contains(path, TempDir) || strings.Contains(path, ExecutableName) || strings.Contains(path, ".git") || strings.Contains(path, "output") {
		return true
	}

	return false
}

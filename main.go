package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
	TODO
		Check for args
		Cross-Platform directory args
		Ignore desired files / folders

*/

func main() {

	// TODO Build output for other systems aswell

	// Find name of current executable
	execPath := strings.Split(os.Args[0], string(filepath.Separator))
	ExecutableName = execPath[len(execPath)-1]

	fmt.Println(os.Args[0])

	fmt.Println("Creating temporary directory")
	err := os.Mkdir(TempDir, 0744)
	printError(err)

	fmt.Print("Implementing files into the binary\n\n")
	err = ImplementFiles(".")
	printError(err)

	fmt.Println("\nImplementing the packr library")
	err = ImplementPackrLibrary()
	printError(err)

	fmt.Println("Implementing installer files")
	err = ImplementInstallerFiles()
	printError(err)

	fmt.Println("Building new executable")
	err = BuildNewBinary("windows")
	printError(err)

	err = CleanUp()
	printError(err)
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

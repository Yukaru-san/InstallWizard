package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {

	// Interrupt handler
	SetupCloseHandler()

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

	fmt.Println("\nTelling your installer it's name")
	err = ImplementInstallerName()
	printError(err)

	fmt.Println("Implementing required libraries")
	err = ImplementPackrLibrary()
	printError(err)
	err = ImplementSqweekLibrary()
	printError(err)

	fmt.Println("Implementing installer files")
	err = ImplementInstallerFiles()
	printError(err)

	fmt.Println("Building new executables")
	err = BuildNewBinary()
	printError(err)

	err = CleanUp()
	printError(err)
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// SetupCloseHandler catches an interrupt signal from the host pc and cleans the files before exiting
func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Program interrupted. Cleaning up and exiting.")
		CleanUp()
		os.Exit(0)
	}()
}

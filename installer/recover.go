// +build !linux,windows darwin

package main

import (
	"os"

	"github.com/gen2brain/dlgs"
)

// RecoverFileStructure recovers files into DataStruct
func RecoverFileStructure() {

	// Installation path
	desiredDirectory := SelectPath()

	// Unpack and create files
	CreateFiles(desiredDirectory)

	// Done!
	dlgs.Info("InstallWizard", ProgramName+" has been installed. You are ready to go!")
}

// SelectPath returns the path the user wishes to install the programm into
func SelectPath() string {
	dlgs.Info("InstallWizard", "This will install "+ProgramName+" onto your harddrive. Please select your desired installation folder")

	dir, _, err := dlgs.File("Please select the directory you wish your program to be installed in", "", true)

	if err != nil {
		dlgs.Info("InstallWizard", "Setup successfully cancled.")
		os.Exit(0)
	}

	return dir
}

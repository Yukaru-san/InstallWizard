// +build linux,!windows,!darwin

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RecoverFileStructure recovers files into DataStruct
func RecoverFileStructure() {

	// Installation path
	desiredDirectory := SelectPath()

	fmt.Println("Unpacking files...")

	// Unpack and create files
	CreateFiles(desiredDirectory)

	fmt.Println("Finished. You are ready to go!")
}

// SelectPath returns the path the user wishes to install the programm into
func SelectPath() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("This will install " + ProgramName + ".")

	dir := ""

	if len(os.Args) == 1 {
		fmt.Print("\nYou didn't specify an installation path when starting the installer.\nPlease do so now:")

		input, _ := reader.ReadString('\n')
		input = strings.ReplaceAll(input, "\n", "")

		dir = input
	} else {
		dir = os.Args[1]
		fmt.Println("You chose " + dir + " as your installation path.")
	}

	fmt.Print("\n\n")
	return dir
}

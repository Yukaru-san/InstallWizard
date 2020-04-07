package main

import (
	"fmt"
	"os"

	"github.com/Yukaru-san/InstallWizard/explorer"
)

func main() {
	err := os.Mkdir(explorer.TempDir, 744)
	printError(err)

	err = explorer.Explore("")
	printError(err)

	err = explorer.SaveDataAsJSON()
	printError(err)

}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*
fmt.Println("Saved in:", fmt.Sprint(gaw.GetHome(), string(filepath.Separator), ".dmanager", string(filepath.Separator), "TestDirectory"))

	err := Explore("")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		CreateFiles(fmt.Sprint(gaw.GetHome(), string(filepath.Separator), ".dmanager", string(filepath.Separator), "TestDirectory"))
	}
*/

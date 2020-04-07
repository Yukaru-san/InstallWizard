package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir(TempDir, 744)
	printError(err)

	err = Explore("")
	printError(err)

	err = SaveDataAsJSON()
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

package commands

import (
	"fmt"
	"hjbdev/pvm/theme"
	"log"
	"os"
	"path/filepath"
)

func Path() {
	theme.Title("pvm: PHP Version Manager")

	// get home dir
	exPath, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	homeDir := filepath.Dir(exPath)
	fmt.Println("Add the following directory to your PATH:")
	fmt.Println("    " + filepath.Join(homeDir, ".pvm", "bin"))
}

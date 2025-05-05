package commands

import (
	"fmt"
	"hjbdev/pvm/common"
	"hjbdev/pvm/theme"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Use(args []string) {
	threadSafe := true

	if len(args) < 1 {
		theme.Error("You must specify a version to use.")
		return
	}

	if len(args) > 1 {
		if args[1] == "nts" {
			threadSafe = false
		}
	}

	// get home dir
	exPath, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	homeDir := filepath.Dir(exPath)

	// check if .pvm folder exists
	if _, err := os.Stat(filepath.Join(homeDir, ".pvm")); os.IsNotExist(err) {
		theme.Error("No PHP versions installed")
		return
	}

	// check if .pvm/versions folder exists
	if _, err := os.Stat(filepath.Join(homeDir, ".pvm", "versions")); os.IsNotExist(err) {
		theme.Error("No PHP versions installed")
		return
	}

	// check if .pvm/bin folder exists
	binPath := homeDir
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		os.Mkdir(binPath, 0755)
	}

	// get all folders in .pvm/versions
	versions, err := os.ReadDir(filepath.Join(homeDir, ".pvm", "versions"))
	if err != nil {
		log.Fatalln(err)
	}

	var selectedVersion *versionMeta
	// loop over all found installed versions
	for i, version := range versions {
		safe := true
		if strings.Contains(version.Name(), "nts") || strings.Contains(version.Name(), "NTS") {
			safe = false
		}
		foundVersion := common.ComputeVersion(version.Name(), safe, "")
		if threadSafe == foundVersion.ThreadSafe && strings.HasPrefix(foundVersion.String(), args[0]) {
			selectedVersion = &versionMeta{
				number: foundVersion,
				folder: versions[i],
			}
		}
	}

	if selectedVersion == nil {
		theme.Error("The specified version is not installed.")
		return
	}

	requestedVersion := common.ComputeVersion(args[0], threadSafe, "")
	if requestedVersion.Minor == -1 {
		theme.Warning(fmt.Sprintf("No minor version specified, assumed newest minor version %s.", selectedVersion.number.String()))
	} else if requestedVersion.Patch == -1 {
		theme.Warning(fmt.Sprintf("No patch version specified, assumed newest patch version %s.", selectedVersion.number.String()))
	}

	versionFolderPath := filepath.Join(homeDir, ".pvm", "versions", selectedVersion.folder.Name())
	
	// create directory link to ext directory
	extensionDirPath := versionFolderPath
	extensionLinkPath := filepath.Join(binPath, "current")

	// delete the old link first if it exists
	if _, err := os.Stat(extensionLinkPath); err == nil {
		cmd := exec.Command("cmd", "/C", "rmdir", extensionLinkPath)
		_, err := cmd.Output()
		if err != nil {
			log.Fatalln("Error deleting ext directory directory link:", err)
			return
		}
	}

	// create directory link - uses cmd since using os.Symlink did require extra permissions
	cmd := exec.Command("cmd", "/C", "mklink", "/J", extensionLinkPath, extensionDirPath)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalln("Error creating ext directory symlink:", err)
		return
	} else {
		theme.Info(string(output))
	}
	// end of ext directory link creation

	theme.Success(fmt.Sprintf("Using PHP %s", selectedVersion.number))
}

type versionMeta struct {
	number common.Version
	folder os.DirEntry
}

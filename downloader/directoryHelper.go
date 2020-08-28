package downloader

import (
	"fmt"
	"log"
	"os"
)

// Our interface for the DirectoryHelper
type IDirectoryHelper interface {
	CreateRequiredFolders(sub string)
}

// Our concrete implementation of the DirectoryHelper
type CDirectoryHelper struct{}

// Our provider for the DirectoryHelper
func InitDirectoryHelper() IDirectoryHelper {
	return CDirectoryHelper{}
}

func (dh CDirectoryHelper) CreateRequiredFolders(sub string) {
	if exists := checkIfDirExists(sub); !exists {
		done := createDir(sub + "/")
		if !done {
			log.Fatal("Failed to create folder")
		}
	}
}

/***************************
     Private Functions
***************************/

func checkIfDirExists(folder string) bool {

	homeDir, _ := os.UserHomeDir()

	path := fmt.Sprintf("%s%s%s", homeDir, "/Documents/reddit-downloader/", folder)

	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func createDir(folder string) bool {

	homeDir, _ := os.UserHomeDir()
	path := fmt.Sprintf("%s%s%s", homeDir, "/Documents/reddit-downloader/", folder)

	if err := os.Mkdir(path, 0755); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

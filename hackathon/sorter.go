package hackathon

import (
	"fmt"
	"os"
	"path/filepath"
)

func sortFiles(rootDir string) {
	// Implement file sorting logic, you can use filepath.Walk to navigate directory and sort files
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// Ignore hidden folders and files
		if info.IsDir() && info.Name()[0] == '.' {
			return filepath.SkipDir
		} else if info.Name()[0] == '.' {
			return nil
		}
		// move the file to a folder based on its extension
		extension := filepath.Ext(info.Name())
		// create a folder with the extension name in root dir if it doesn't exist already
		extensionDir := filepath.Join(rootDir, extension[1:])
		if _, err := os.Stat(extensionDir); os.IsNotExist(err) {
			os.Mkdir(extensionDir, os.ModePerm)
		}
		// move the file to the folder
		err = os.Rename(path, filepath.Join(extensionDir, info.Name()))
		if err != nil {
			fmt.Println(err)
		}
		return nil

	})
	if err != nil {
		fmt.Println(err)
	}
}

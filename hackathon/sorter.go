package hackathon

import (
	"fmt"
	"os"
	"path/filepath"
)

func sortFiles(rootDir string) error {
	// Implement file sorting logic, you can use filepath.Walk to navigate directory and sort files
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// Ignore hidden folders and files
		if filepath.Base(path)[0] == '.' {
			return nil
		}
		if info.IsDir() {
			// Ignore directories
			return nil
		}
		// Get the file extension
		ext := filepath.Ext(path)
		// Create a folder with the file extension
		dir := filepath.Join(rootDir, ext[1:])
		// Create the folder if it doesn't exist already
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, 0755)
		} else if err != nil {
			return err
		}
		// Move the file to the folder
		newPath := filepath.Join(dir, filepath.Base(path))
		err = os.Rename(path, newPath)
		if err != nil {
			return err
		}
		fmt.Printf("Moved %s to %s\n", path, newPath)
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("Files sorted successfully")
	return nil
}

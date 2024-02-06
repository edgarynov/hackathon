package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Replace "/" with the path of the directory you want to sort
	rootDir := "\\wsl.localhost/Ubuntu-20.04/home/christian/Documents/"

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Ignore directories and hidden files or folders
		if path == rootDir || info.IsDir() || string(path[len(path)-1]) == "~" {
			return nil
		}
		// Get the extension of the current file
		extension := filepath.Ext(info.Name())
		folderPath := filepath.Join(rootDir, "Documents") + string(os.PathSeparator)

		// Set the destination directory based on the extension
		switch extension {
		case ".pdf":
			folderPath += "PDF"
		case ".docx", ".doc":
			folderPath += "Word"
		case ".xlsx", ".xls":
			folderPath += "Excel"
		case ".pptx", ".ppt":
			folderPath += "PowerPoint"
		case ".txt":
			folderPath += "Text"
		case ".jpg", ".jpeg", ".png", ".gif":
			folderPath += "Images"
		case ".mp3", ".wav", ".flac":
			folderPath += "Audio"
		case ".mp4", ".avi", ".mkv":
			folderPath += "Video"
		default:
			folderPath += "Others"
		}

		// Create the destination directory if it doesn't exist yet
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			_ = os.MkdirAll(folderPath, os.ModePerm)
		}

		// Create the destination path for the current file
		destPath := strings.ReplaceAll(strings.TrimPrefix(path, rootDir), "\\", "/")
		destFile := folderPath + string(os.PathSeparator) + filepath.Base(destPath)
		// Copy the file to its new location
		err = os.Rename(path, destFile)
		if err != nil {
			fmt.Println(err)
			return err
		} else {
			fmt.Printf("File sorted: %s -> %s\n", path, destFile)
		}

		// Move the file to its corresponding folder
		err = os.Rename(path, folderPath+info.Name())
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		return nil

	})
	// Check if there was an error while walking the directory
	if err != nil {
		fmt.Println("Error:", err)

	} else {
		fmt.Println("Files sorted successfully")
	}
}

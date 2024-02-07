package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func sort() {
	sourceDir := "chemin début"
	destinationDir := "chemin arriver"

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("erreur nav fichier : %s\n", err)
			return err
		}

		if !info.IsDir() {
			moveFile(path, info, destinationDir)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("erreur nav fichier : %s\n", err)
	}
}

func moveFile(filePath string, fileInfo os.FileInfo, destinationDir string) {
	fileName := fileInfo.Name()
	ext := filepath.Ext(fileName)
	destDir := filepath.Join(destinationDir, strings.TrimPrefix(ext, "."))

	//Erzeugt den Zielordner, falls er noch nicht existiert.
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.MkdirAll(destDir, 0755)
		if err != nil {
			fmt.Printf("erreur création fichier arriver : %s\n", err)
			return
		}
	}

	destFilePath := filepath.Join(destDir, fileName)

	err := os.Rename(filePath, destFilePath)
	if err != nil {
		fmt.Printf("erreurs lors deplacement fichier : %s\n", err)
	} else {
		fmt.Printf("déplacement réussi : %s\n", fileName)
	}
}

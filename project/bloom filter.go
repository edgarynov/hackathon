package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/willf/bloom"
)

func bloomfilter() {
	//Startpfad oder der Filter muss die Dateien nehmen, um sie in den Filter zu setzen
	sourceDir := "chemin début"

	//setup du filtre a 500 éléments et 0.01% de chance d'avoir un faux positif
	filter := bloom.NewWithEstimates(500, 0.0001)

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("erreur nav fichier : %s\n", err)
			return err
		}

		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("erreur lecture dossier : %s\n", err)
				return err
			}

			filter.Add(data)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("erreur nav dossier : %s\n", err)
	}
}

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func identity() {
	sourceDir := "chemin début"
	destinationDir := "chemin fin"

	sourceFileHashes := make(map[string]string)
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("erreur navigation file : %s\n", err)
			return err
		}

		if !info.IsDir() {
			sourceFileHashes[path], err = checksum(path)
			if err != nil {
				fmt.Printf("erreur calcul hachage MD5 : %s\n", err)
				return err
			}
		}
		return nil
	})

	filepath.Walk(destinationDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf(" erreur navigation file : %s\n", err)
			return err
		}

		if !info.IsDir() {
			fileHash, err := checksum(path)
			if err != nil {
				fmt.Printf("erreur calcul hachage MD5 : %s\n", err)
				return err
			}

			if sourceHash, ok := sourceFileHashes[path]; ok && sourceHash == fileHash {
				err := os.Remove(path)
				if err != nil {
					fmt.Printf("erreur suppression doublon : %s\n", err)
				} else {
					fmt.Printf("doublon : %s\n", path)
				}
			}
		}
		return nil
	})
}

func checksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

//Ich wusste nicht, ob ich einen fünften Satz auf Deutsch schreiben sollte, also wird sie hier sein.

package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func IsFileExist(fileName string) bool {
	dir := "./database"

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		slice := strings.Split(file.Name(), "_")
		if slice[0] == fileName {
			return true
		}
	}

	return false
}

func GetFileExist(fileName string) string {
	dir := "./database"

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		slice := strings.Split(file.Name(), "_")
		if slice[0] == fileName {
			return fmt.Sprintf("%s/%s", dir, file.Name())
		}
	}

	return ""
}

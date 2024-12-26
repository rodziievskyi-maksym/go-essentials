package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func createEmptyFile() *os.File {
	pathToStorage := "storage/"
	currentTime := time.Now().Format("2006-01-02T15:04:05")
	format := ".txt"
	createDirIfNotExists(pathToStorage)

	fileName := fmt.Sprintf("%s%s%s", pathToStorage, currentTime, format)
	myFile, err := os.Create(fileName)
	check(err)
	log.Println("Empty file create successfully.", fileName)

	return myFile
}

func writeIntoFile(data string, file *os.File) {
	_, err := file.WriteString(data)
	check(err)
}

func createDirIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

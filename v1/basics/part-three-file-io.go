package basics

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

/*
	Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.

	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
*/

/*
----------	0000	no permissions
-rwx------	0700	read, write, & execute only for owner
-rwxrwx---	0770	read, write, & execute for owner and group
-rwxrwxrwx	0777	read, write, & execute for owner, group and others
---x--x--x	0111	execute
--w--w--w-	0222	write
--wx-wx-wx	0333	write & execute
-r--r--r--	0444	read
-r-xr-xr-x	0555	read & execute
-rw-rw-rw-	0666	read & write
-rwxr-----	0740	owner can read, write, & execute; group can only read; others have no permissions
*/

func WorkWithFiles() {

	inputData := inputData()
	file := createEmptyFile()
	fileInfoFunc(file.Name())
	writeIntoFile(inputData, file)

	readOutOfFile(file.Name())

	defer file.Close()

	//readWholeFile(fileValue)
	//readSomeBytesFromFile(fileValue, 10)
}

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

func inputData() string {
	pl("Add a new word: ")
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadBytes('\n')
	check(err)

	return string(data)
}

func fileInfoFunc(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size ", fileInfo.Size(), " bytes")
	fmt.Println("Permissions:", fileInfo.Mode()) //Permissions: -rw-r--r--
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
}

func readOutOfFile(filePath string) {
	file, err := os.Open(filePath)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pl(scanner.Text())
	}

}

func readWholeFile(fileName string) {
	file, err := os.ReadFile(fileName)
	check(err)
	pl("Read whole file: ", string(file))
}

func readSomeBytesFromFile(fileName string, mountOfBytes int) {
	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	fileValue := make([]byte, mountOfBytes)
	numberOfBytes, err := file.Read(fileValue)
	check(err)
	fmt.Printf("%d bytes: %s\n", numberOfBytes, string(fileValue[:numberOfBytes]))
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

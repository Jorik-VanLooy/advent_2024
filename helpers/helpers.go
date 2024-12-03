package helpers

import (
	"fmt"
	"io"
	"log"
	"os"
)

func OpenFile(file string) []byte {
	// using os.Open() function to open file
	fileData, errOpeningFile := os.Open(file)
	if errOpeningFile != nil {
		fmt.Printf("failed to open file due to %s", errOpeningFile)
	}
	// close the file to ensure no memory leak
	defer fileData.Close()
	// get the content using ioutil.ReadAll() package returns data in form of bytes 40 50 ..
	data, errReadingFileData := io.ReadAll(fileData)
	if errReadingFileData != nil {
		log.Fatalf("Sorry could not read file content due to %s", errReadingFileData)
	}

	return data
}

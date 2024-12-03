package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData() ([]float64, []float64) {
	file := "input_1.txt"
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
	// convert the content to string for humanreadable format i.e Hello
	oneString := strings.Replace(string(data), "\n", "   ", -1)
	strArr := strings.Split(oneString, "   ")
	var leftArr []float64
	var rightArr []float64

	for index := range strArr {
		number, _ := strconv.ParseFloat(strArr[index], 64)
		if index%2 == 0 {
			leftArr = append(leftArr, number)
		} else {
			rightArr = append(rightArr, number)
		}
	}
	sort.Float64s(leftArr)
	sort.Float64s(rightArr)

	return leftArr, rightArr
}

func DayOneChallengeOne() {
	leftArr, rightArr := getData()
	var distances []float64
	for index := range leftArr {
		distances = append(distances, math.Abs(leftArr[index]-rightArr[index]))
	}

	out := int(distances[0])
	for _, number := range distances {
		out += int(number)
	}
	fmt.Println(out)
}

func DayOneChallengeTwo() {
	var simScore []int
	leftArr, rightArr := getData()
	for _, leftNumber := range leftArr {
		count := 0
		for _, rightNumber := range rightArr {
			if leftNumber == rightNumber {
				count += 1
			}
		}
		simScore = append(simScore, count*int(leftNumber))
	}

	out := simScore[0]

	for _, number := range simScore {
		out += number
	}

	fmt.Println(out)
}

func main() {
	challenge := os.Args[1]
	switch challenge {
	case "1":
		DayOneChallengeOne()
	case "2":
		DayOneChallengeTwo()
	}
}

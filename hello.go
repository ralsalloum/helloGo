package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

var sumValue int = 0

func countFileValues(file string) {

	defer wg.Done()

	//fmt.Println(file)

	openedfile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(openedfile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		intValue, err := strconv.Atoi(fileScanner.Text())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sumValue = sumValue + intValue
	}

	//fmt.Println(sumValue)
}

func main() {
	var files []string = []string{"textFile1.text", "textFile2.text", "textFile3.text", "textFile4.text"}

	currentWorkingDir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	for i, file := range files {
		wg.Add(1)

		go countFileValues(currentWorkingDir + "\\files\\" + file)

		fmt.Println("File #" + strconv.Itoa(i) + " is read")
	}

	wg.Wait()

	fmt.Println("Sum value = " + strconv.Itoa(sumValue))
}

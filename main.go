package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func checkError(err error){

	if err != nil{
		panic(err)
	}
}
func main()  {
	
	file, err := os.Open("rg-for-range.txt")
	checkError(err)

	defer file.Close()

	re, err := regexp.Compile(`(.*)\.go:\d+:`)
	checkError(err)

	scanner := bufio.NewScanner(file)

    filesMap := make(map[string]int)

	for scanner.Scan() {
		line :=scanner.Text()

		matches := re.FindStringSubmatch(line)
		if matches == nil {
			panic(errors.New("no match"))
		}

		fileName := matches[1] + ".go"

		_, ok := filesMap[fileName]

		if(ok){
			filesMap[fileName]++
		}else{
			filesMap[fileName] = 1
		}

	}

	for fileName, count := range filesMap {
        fmt.Printf("%s: %d\n", fileName, count)
    }

	fmt.Printf("Unique files: %d\n", len(filesMap))

	scannerErr := scanner.Err()

	checkError(scannerErr)
}
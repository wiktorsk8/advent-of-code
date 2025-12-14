package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func loadIds() [][]int {
	input, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var ids [][]int

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)

		ranges := strings.Split(line, ",")

		for _, element := range ranges {
			splittedRange := strings.Split(element, "-")
			var idsRange []int

			for _, id := range splittedRange {
				idInt, err := strconv.Atoi(strings.TrimSpace(id))

				if err != nil {
					log.Fatal(err.Error())
				}

				idsRange = append(idsRange, idInt)
			}

			ids = append(ids, idsRange)
		}
	}

	return ids
}

func detectInvalidIds(start, end int) []int {
	var invalidIds []int

	for i := start; i <= end; i++ {

		iStr := strconv.Itoa(i)

		if len(iStr)%2 != 0 {
			continue
		}

		if isMadeFromDuplicateSequence(iStr) {
			invalidIds = append(invalidIds, i)
		}

	}

	return invalidIds

}

func isMadeFromDuplicateSequence(id string) bool {
	size := len(id)
	mid := size / 2

	return id[:mid] == id[mid:]
}

func main() {
	ids := loadIds()

	resultsChan := make(chan int)
	var wg sync.WaitGroup

	for _, idRange := range ids {
		wg.Add(1)

		go func(start, end int) {
			defer wg.Done()
			partialSum := 0
			invalidIds := detectInvalidIds(start, end)
			for _, id := range invalidIds {
				partialSum += id
			}

			resultsChan <- partialSum
		}(idRange[0], idRange[1])
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	totalSum := 0

	for sum := range resultsChan {
		totalSum += sum
	}

	fmt.Printf("hello %d \n", totalSum)
}

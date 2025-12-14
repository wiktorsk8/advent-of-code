package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	addedUpIds := 0
	for _, idRange := range ids {
		invalidIdsInRange := detectInvalidIds(idRange[0], idRange[1])
		for _, id := range invalidIdsInRange {
			addedUpIds += id
		}
	}

	fmt.Printf("hello %d", addedUpIds)
}

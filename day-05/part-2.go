package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type MapEntry struct {
	destinationRangeStart uint64
	sourceRangeStart      uint64
	sourceRangeEnd     uint64
}

type MapRange struct {
	low uint64
	high uint64
}

func readFile(filename string) []string {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	return lines
}

func readInput() []string {
	return readFile("input.txt")
}

func readTestInput() []string {
	return readFile("input-test.txt")
}

func main() {
	input := readInput()

	seeds := []uint64{}

	var mapAssigned string

	maps := make(map[string][]MapEntry)
	mapRanges := make(map[string]MapRange)

	mapOrder := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for idx, val := range input {
		// Ingest seed numbers
		if idx == 0 {
			numStrs := strings.Split(val[strings.Index(val, ":")+2:], " ")
			for _, numStr := range numStrs {
				num, _ := strconv.ParseUint(numStr, 10, 64)
				seeds = append(seeds, num)
			}

			continue
		}

		// Skip blank lines
		if val == "" {
			continue
		}

		// Assign map to be ingested
		if _, err := strconv.Atoi(string(val[0])); err != nil {
			mapAssigned = strings.TrimSuffix(val, " map:")
			mapRanges[mapAssigned] = MapRange{ uint64(math.MaxUint64), 0}

			continue
		}

		// Process numbers
		nums := []uint64{}
		for _, numStr := range strings.Split(val, " ") {
			num, _ := strconv.ParseUint(numStr, 10, 64)
			nums = append(nums, num)
		}
		maps[mapAssigned] = append(maps[mapAssigned], MapEntry{nums[0], nums[1], nums[1] + nums[2]})

	}

	for key, _ := range(maps) {
		slices.SortFunc(maps[key], func(a, b MapEntry) int {
			return cmp.Compare(a.sourceRangeStart, b.sourceRangeStart)
		})
	}

	minLocation := uint64(math.MaxUint64)
	timeStart := time.Now()
	for seedGroup := 0; seedGroup < len(seeds); seedGroup += 2 {
		start := seeds[seedGroup]
		end := start + seeds[seedGroup + 1]
		for seedNumber := start; seedNumber < end; seedNumber++ {
			seed := seedNumber

			for _, mapName := range mapOrder {

				rangeMapEntry := MapEntry{}
				rangeMapEntryFound := false

				for _, mapEntry := range maps[mapName] {
					if mapEntry.sourceRangeStart <= seed  {
						rangeMapEntry = mapEntry
						rangeMapEntryFound = true
					}
				}
				
				// for i, j := 0, len(maps[mapName]) - 1; i < j; i, j = i + 1, j - 1  {
				// 	if maps[mapName][j].sourceRangeStart <= seed {
				// 		rangeMapEntry = maps[mapName][j]
				// 		rangeMapEntryFound = true
				// 		break
				// 	}

				// 	if maps[mapName][i].sourceRangeStart > seed {
				// 		break
				// 	}

				// 	rangeMapEntry = maps[mapName][i]
				// 	rangeMapEntryFound = true
				// }



				

				if rangeMapEntryFound && rangeMapEntry.sourceRangeEnd > seed{
					seed = rangeMapEntry.destinationRangeStart + seed - rangeMapEntry.sourceRangeStart
				}
				// fmt.Println(seed, rangeMapEntry, mapName, maps[mapName])
			}

			minLocation = min(seed, minLocation)
			// break
		}

		break
	}
	timeEnd := time.Now()
	fmt.Println("finised in ", timeEnd.Sub(timeStart).String())
	fmt.Println(minLocation)
	fmt.Println("done")
}

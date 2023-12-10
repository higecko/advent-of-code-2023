package main

import (
	"fmt"
	"os"
	"strconv"
	"slices"
	"strings"
)

type MapEntry struct {
	destinationRangeStart uint64
	sourceRangeStart uint64
	sourceRangeLength uint64
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

  mapOrder := []string {
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
	  "humidity-to-location",
	}

	for idx, val := range(input) {
		// Ingest seed numbers
		if idx == 0 {
			numStrs := strings.Split(val[strings.Index(val, ":") + 2:], " ")
			for _, numStr := range(numStrs) {
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
			continue
		}

		// Process numbers
		nums := []uint64{}
		for _, numStr := range(strings.Split(val, " ")) {
			 num, _ := strconv.ParseUint(numStr, 10, 64)
			 nums = append(nums, num)
		}
		maps[mapAssigned] = append(maps[mapAssigned], MapEntry{nums[0], nums[1], nums[2]})
	}

	seedLocations := []uint64{}
  for _, seed := range(seeds) {
    for _, mapName := range(mapOrder) {
      rangeMapEntry := MapEntry{}
      rangeMapEntryFound := false

      for _, mapEntry := range(maps[mapName]){
        if mapEntry.sourceRangeStart <= seed && ((!rangeMapEntryFound && mapEntry.sourceRangeStart == 0)|| mapEntry.sourceRangeStart > rangeMapEntry.sourceRangeStart) {
          rangeMapEntry = mapEntry
          rangeMapEntryFound = true
        }
      }

      if rangeMapEntryFound && rangeMapEntry.sourceRangeStart + rangeMapEntry.sourceRangeLength > seed {
        seed = rangeMapEntry.destinationRangeStart + seed - rangeMapEntry.sourceRangeStart
      }
      // fmt.Println(seed, rangeMapEntry, mapName, maps[mapName])
    }

    seedLocations = append(seedLocations, seed)
  }
 
  fmt.Println(slices.Min(seedLocations))
}
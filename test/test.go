package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
func main() {
	// date_string := "06/07/2024 12:21:33"
	// date, _ := time.Parse("02/01/2006 15:04:05", date_string)
	fmt.Print("\n")

}

package main

import (
	"encoding/csv"
	"log"
	"os"
)

//Create csv meta-data struct?

//type csv_metadata struct {}


// Csv file reader from Go packages
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

// Assuming header row in csv

func getNumRecordsHeader(csv [][]string) int {
	csv_file := csv
	return len(csv_file) - 1
}

package main

import (
	"fmt"
)

func main() {
	records := readCsvFile("username.csv")
	fmt.Println(records)
	fmt.Printf("The number of fields is %d", getFieldNum(records))
	fmt.Println("csv[0] is ", records[0])
	fmt.Println("the len of csv[0] is", len(records[0]))
}

// Function that returns number of fields
func getFieldNum(csv [][]string) int {
	return len(csv[0][0])
}

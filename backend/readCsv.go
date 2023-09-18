package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Powerlifter struct {
	Age       int
	Squat1Kg  float64
	Bench1Kg  float64
}


func readCsv() map[string][]interface{} {
	// Open the CSV file for reading
	file, err := os.Open("./openpowerlifting.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Create maps to store the data
	data := map[string][]interface{}{
		"age":   nil,
		"squat": nil,
		"bench": nil,
	}

	// Read and process each row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Parse the relevant columns
		age := parseInt(record[4])
		squat1Kg := parseFloat(record[9])
		bench1Kg := parseFloat(record[16])

		// Check if age is positive and stats are non-negative
		if age > 0 && squat1Kg > 0 && bench1Kg > 0 {
			// Append data to the corresponding slice in the map
			data["age"] = append(data["age"], age)
			data["squat"] = append(data["squat"], squat1Kg)
			data["bench"] = append(data["bench"], bench1Kg)
		}
	}

	return data
}

func parseInt(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return 0 // Default to 0 if parsing fails
	}
	return result
}

func parseFloat(s string) float64 {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	if err != nil {
		return 0.0 // Default to 0.0 if parsing fails
	}
	return result
}

package game

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCSVFile(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to open the CSV file: %s\n", err))
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	return reader.ReadAll()
}

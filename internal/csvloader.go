package internal

import (
	"encoding/csv"
	"get-item-json-service/internal/handler"
	"get-item-json-service/model"
	"io"
	"log"
	"os"
	"strconv"
)

func (h *handler.Handler) LoadEmployeesFromCSV() ([]model.Employee, error) {
	file, err := os.Open(h.CSVFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var employees []model.Employee

	_, err = reader.Read()
	if err != nil {
		log.Printf("Error reading CSV header: %v\n", err)
		return nil, err
	}

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading CSV row: %v\n", err)
			return nil, err
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Printf("Error parsing ID: %v\n", err)
			continue
		}

		employee := model.Employee{
			ID:         id,
			UID:        row[1],
			CN:         row[2],
			Department: row[3],
			Title:      row[4],
			Who:        row[5],
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

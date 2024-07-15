package handler

import (
	"encoding/csv"
	"get-item-json-service/model"
	"io"
	"log"
	"os"
	"strconv"
)

type CsvManager interface {
	Open(name string) (*os.File, error)
	NewReader(r io.Reader) *csv.Reader
}

func (h *Handler) LoadEmployeesFromCSV() (map[int]model.Employee, error) {
	file, err := h.CsvManager.Open(EmployeeCSVFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := h.CsvManager.NewReader(file)

	employees := make(map[int]model.Employee)

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

		employees[id] = employee
	}

	return employees, nil
}

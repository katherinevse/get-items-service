package handler

import (
	"encoding/json"
	"get-item-json-service/model"
	"net/http"
	"strconv"
	"strings"
)

const (
	EmployeeCSVFile = "ueba.csv"
)

type Handler struct {
	CsvManager CsvManager
}

func New(CsvManager CsvManager) *Handler {
	return &Handler{CsvManager: CsvManager}
}

// GetEmployeesHandler GEt/get-items?id=?
func (h *Handler) GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query().Get("id")
	if queryValues == "" {
		http.Error(w, "Missing query parameter: id", http.StatusBadRequest)
		return
	}
	idStrings := strings.Split(queryValues, ",")

	if len(idStrings) == 0 {
		http.Error(w, "No IDs provided", http.StatusBadRequest)
		return
	}

	employees, err := h.LoadEmployeesFromCSV()
	if err != nil {
		http.Error(w, "Failed to load employees from CSV", http.StatusInternalServerError)
		return
	}

	employeeResults := make([]model.Employee, 0, len(idStrings))

	for _, idStr := range idStrings {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}

		if emp, exists := employees[id]; exists {
			employeeResults = append(employeeResults, emp)
		}
	}

	jsonData, err := json.Marshal(employeeResults)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

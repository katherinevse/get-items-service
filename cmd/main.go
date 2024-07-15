package main

import (
	"get-item-json-service/internal/csv_manager"
	"get-item-json-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	csvManager := &csv_manager.CsvManager{}
	newHandler := handler.New(csvManager)

	http.HandleFunc("/get-items", newHandler.GetEmployeesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

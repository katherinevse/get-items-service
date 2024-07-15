package main

import (
	"get-item-json-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	handler := handler.NewHandler(handler.EmployeeCSVFile)

	http.HandleFunc("/get-items", handler.GetEmployeesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

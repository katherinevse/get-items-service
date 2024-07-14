package main

import (
	"get-item-json-service/csvloader"
	"log"
	"net/http"
)

func main() {
	handler := csvloader.NewHandler(csvloader.EmployeeCSVFile)

	http.HandleFunc("/get-items", handler.GetEmployeesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"get-item-json-service/internal"
	"log"
	"net/http"
)

func main() {
	handler := internal.NewHandler(internal.EmployeeCSVFile)

	http.HandleFunc("/get-items", handler.GetEmployeesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

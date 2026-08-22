package main

import "fmt"

type Employee struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	// Pascal Case
	// Eg. CalculateArea, UserInfo, NewHTTPRequest

	// Snake Case
	// Eg. user_id, first_name, http_request

	// Upper Case 
	// Use Case: Constants

	// Mixed Case
	// Eg. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeID:", employeeID)
}
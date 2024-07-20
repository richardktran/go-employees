package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type employee struct {
	ID           int    `json:"id"`
	Name         string `json:"employee_name"`
	Age          int    `json:"employee_age"`
	Salary       int    `json:"employee_salary"`
	ProfileImage string `json:"profile_image"`
}

type employeesResponse struct {
	Status string     `json:"status"`
	Data   []employee `json:"data"`
}

func main() {
	resp, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
	}

	file, err := os.OpenFile("data/employees.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// debug to see resp body
	var employees employeesResponse

	if err := json.NewDecoder(resp.Body).Decode(&employees); err != nil {
		panic(err)
	}

	file.WriteString("id,name,salary,age,profile_image\n")

	for _, employee := range employees.Data {
		newEmployee := fmt.Sprintf("%d,%s,%d,%d,%s\n", employee.ID, employee.Name, employee.Salary, employee.Age, employee.ProfileImage)
		file.WriteString(newEmployee)
	}
}

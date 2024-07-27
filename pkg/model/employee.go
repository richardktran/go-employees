package model

type Employee struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Age          int    `json:"age"`
	ProfileImage string `json:"profile_image"`
}

type EmployeeCreation struct {
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Age          int    `json:"age"`
	ProfileImage string `json:"profile_image"`
}

type EmployeeUpdate struct {
	Name         string `json:"name"`
	Salary       int    `json:"salary"`
	Age          int    `json:"age"`
	ProfileImage string `json:"profile_image"`
}

type EmployeeResponse struct {
	Status string     `json:"status"`
	Data   []Employee `json:"data"`
}

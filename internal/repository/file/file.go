package file

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/richardktran/go-employees/pkg/model"
)

type Repository struct {
	mu       sync.RWMutex
	filePath string
}

func New(filePath string) *Repository {
	return &Repository{filePath: filePath}
}

func (r *Repository) GetEmployees(_ context.Context) ([]model.Employee, error) {
	var employees []model.Employee

	r.mu.RLock()

	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	headerLine := scanner.Text()
	headers := strings.Split(headerLine, ",")

	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	for scanner.Scan() {
		employee := scanner.Text()
		// Split each line by comma
		employeeData := strings.Split(employee, ",")

		id, err := strconv.Atoi(employeeData[headerMap["id"]])
		if err != nil {
			return nil, err
		}

		age, err := strconv.Atoi(employeeData[headerMap["age"]])
		if err != nil {
			return nil, err
		}

		salary, err := strconv.Atoi(employeeData[headerMap["salary"]])
		if err != nil {
			return nil, err
		}

		employees = append(employees, model.Employee{
			ID:           id,
			Name:         employeeData[headerMap["name"]],
			Salary:       salary,
			Age:          age,
			ProfileImage: employeeData[headerMap["profile_image"]],
		})
	}

	defer r.mu.RUnlock()

	return employees, nil
}

func (r *Repository) AddEmployee(ctx context.Context, employee model.EmployeeCreation) error {
	log.Printf("Adding employee to file: %+v", employee)
	ID := r.generateId(ctx)
	r.mu.Lock()
	defer r.mu.Unlock()
	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	log.Printf("Generated ID: %d", ID)

	newEmployee := fmt.Sprintf("%d,%s,%d,%d,%s\n", ID, employee.Name, employee.Salary, employee.Age, employee.ProfileImage)

	log.Printf("New employee: %s", newEmployee)
	_, err = file.WriteString(newEmployee)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) generateId(ctx context.Context) int {
	employees, err := r.GetEmployees(ctx)
	if err != nil {
		return 1
	}

	if len(employees) == 0 {
		return 1
	}

	return employees[len(employees)-1].ID + 1
}

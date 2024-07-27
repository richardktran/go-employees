package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/richardktran/go-employees/configs/env"
	"github.com/richardktran/go-employees/pkg/model"
)

type Repository struct {
	db *sql.DB
}

func New() (*Repository, error) {
	dbHost := env.GET("DB_HOST")
	dbPort := env.GET("DB_PORT")
	dbUser := env.GET("DB_USERNAME")
	dbPass := env.GET("DB_PASSWORD")
	dbName := env.GET("DB_DATABASE")
	dbOptions := url.Values{
		"charset":   {"utf8mb4"},
		"parseTime": {"True"},
		"loc":       {"Local"},
	}
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	connStr := fmt.Sprintf("%s?%s", connection, dbOptions.Encode())
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) GetEmployees(_ context.Context) ([]model.Employee, error) {
	log.Println("Getting employees from MySQL")
	rows, err := r.db.Query("SELECT id, name, salary, age, profile_image FROM employees")
	if err != nil {
		log.Println("Error getting employees from MySQL", err)
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age, &employee.ProfileImage)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *Repository) AddEmployee(_ context.Context, employee model.EmployeeCreation) error {
	_, err := r.db.Exec("INSERT INTO employees (name, age, salary) VALUES (?, ?, ?)", employee.Name, employee.Age, employee.Salary)
	return err
}

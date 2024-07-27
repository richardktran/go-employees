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

func (r *Repository) GetEmployees(_ context.Context, search string, paging *model.Paging) ([]model.Employee, error) {
	offset := (paging.Page - 1) * paging.Limit
	var rows *sql.Rows
	var err error

	if search != "" {
		rows, err = r.db.
			Query("SELECT id, name, salary, age, profile_image FROM employees WHERE MATCH (name) AGAINST (? IN NATURAL LANGUAGE MODE) LIMIT ? OFFSET ?",
				search,
				paging.Limit,
				offset,
			)
	} else {
		rows, err = r.db.Query("SELECT id, name, salary, age, profile_image FROM employees LIMIT ? OFFSET ?", paging.Limit, offset)
	}

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
	_, err := r.db.Exec("INSERT INTO employees (name, salary, age, profile_image) VALUES (?, ?, ?, ?)", employee.Name, employee.Salary, employee.Age, employee.ProfileImage)
	return err
}

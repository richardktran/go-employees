package employee

import (
	"context"

	"github.com/richardktran/go-employees/pkg/model"
)

type employeeRepository interface {
	GetEmployees(ctx context.Context, search string, paging *model.Paging) ([]model.Employee, error)
	AddEmployee(ctx context.Context, employee model.EmployeeCreation) error
	UpdateEmployee(ctx context.Context, id int, employee model.EmployeeUpdate) error
	DeleteEmployee(_ context.Context, id int) error
}

type Controller struct {
	repo employeeRepository
}

func New(repo employeeRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetEmployees(ctx context.Context, search string, paging *model.Paging) ([]model.Employee, error) {
	return c.repo.GetEmployees(ctx, search, paging)
}

func (c *Controller) AddEmployee(ctx context.Context, employee model.EmployeeCreation) error {
	return c.repo.AddEmployee(ctx, employee)
}

func (c *Controller) UpdateEmployee(ctx context.Context, id int, employee model.EmployeeUpdate) error {
	return c.repo.UpdateEmployee(ctx, id, employee)
}

func (c *Controller) DeleteEmployee(ctx context.Context, id int) error {
	return c.repo.DeleteEmployee(ctx, id)
}

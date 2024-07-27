package employee

import (
	"context"

	"github.com/richardktran/go-employees/pkg/model"
)

type employeeRepository interface {
	GetEmployees(ctx context.Context, paging *model.Paging) ([]model.Employee, error)
	AddEmployee(ctx context.Context, employee model.EmployeeCreation) error
}

type Controller struct {
	repo employeeRepository
}

func New(repo employeeRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetEmployees(ctx context.Context, paging *model.Paging) ([]model.Employee, error) {
	return c.repo.GetEmployees(ctx, paging)
}

func (c *Controller) AddEmployee(ctx context.Context, employee model.EmployeeCreation) error {
	return c.repo.AddEmployee(ctx, employee)
}

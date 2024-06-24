package repositories

import (
	"context"

	"github.com/fazriridwan19/service-employee/domain/models"
	"github.com/jmoiron/sqlx"
)

type IEmployeeRepository interface {
	Create(ctx context.Context, employee *models.Employee) error
	FindById(ctx context.Context, employeeId int64) (models.Employee, error)
	Find(ctx context.Context) ([]models.Employee, error)
	Update(ctx context.Context, employeeId int64, employee models.Employee)
}

type EmployeeRepository struct {
	db sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) IEmployeeRepository {
	return &EmployeeRepository{
		db: *db,
	}
}

// Create implements IEmployeeRepository.
var create = `INSERT INTO employee (name, email) VALUES (?, ?)`

func (e *EmployeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	err := e.db.QueryRowContext(ctx, create, employee.Name, employee.Email).Scan(&employee.Id)
	return err
}

// Find implements IEmployeeRepository.
func (e *EmployeeRepository) Find(ctx context.Context) ([]models.Employee, error) {
	panic("unimplemented")
}

// FindById implements IEmployeeRepository.
func (e *EmployeeRepository) FindById(ctx context.Context, employeeId int64) (models.Employee, error) {
	panic("unimplemented")
}

// Update implements IEmployeeRepository.
func (e *EmployeeRepository) Update(ctx context.Context, employeeId int64, employee models.Employee) {
	panic("unimplemented")
}

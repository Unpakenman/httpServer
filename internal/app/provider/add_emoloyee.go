package provider

import (
	"context"
	"fmt"
	pgclient "httpServer/internal/app/client/pg"
	"httpServer/internal/app/provider/models"
)

type CreateAddEmployeeRequest struct {
	RoleId           int64
	SpecializationId int64
	FirstName        string
	LastName         string
	MiddleName       *string
	BirthDate        string
	Phone            string
	Email            string
}

func (p *goExampleDBProvider) AddEmployee(
	ctx context.Context,
	tx pgclient.Transaction,
	data CreateAddEmployeeRequest,
) (models.Employees, error) {
	var employeesData models.Employees
	err := p.conn.NamedGetContext(
		ctx,
		&employeesData,
		"CreateEmployee",
		nil,
		tx,
		data.RoleId,
		data.SpecializationId,
		data.FirstName,
		data.LastName,
		data.MiddleName,
		data.BirthDate,
		data.Phone,
		data.Email)
	if err != nil {
		return employeesData, fmt.Errorf("CreateEmployee query error: %w", err)
	}
	return employeesData, nil
}

package models

type Employees struct {
	EmployeeID       int64   `db:"employee_id"`
	RoleId           int64   `db:"role_id"`
	SpecializationId int64   `db:"specialization_id"`
	FirstName        string  `db:"first_name"`
	LastName         string  `db:"last_name"`
	MiddleName       *string `db:"middle_name"`
	BirthDate        string  `db:"birthdate"`
	Phone            string  `db:"phone"`
	Email            string  `db:"email"`
	HireDate         string  `db:"hire_date"`
}

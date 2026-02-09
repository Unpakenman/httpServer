package clinics

type AddEmployeeRequest struct {
	RoleId           int64
	SpecializationId int64
	FirstName        string
	LastName         string
	MiddleName       *string
	BirthDate        string
	Phone            string
	Email            string
}

type AddEmployeeResponse struct {
	EmployeeId int64
}

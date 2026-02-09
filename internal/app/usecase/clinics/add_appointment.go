package clinics

type AddAppointmentRequest struct {
	ClinicId        int64
	PatientId       int64
	EmployeeId      int64
	AppointmentDTTM string
	Comment         string
}

type AddAppointmentResponse struct {
	AppointmentId int64
}

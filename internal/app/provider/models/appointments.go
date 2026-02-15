package models

type Appointments struct {
	AppointmentId   int64  `db:"appointment_id"`
	ClinicId        int64  `db:"clinic_id"`
	PatientId       int64  `db:"patient_id"`
	EmployeeId      int64  `db:"employee_id"`
	AppointmentDttm string `db:"appointment_dttm"`
	CreatedAt       string `db:"created_at"`
	Status          string `db:"status"`
	Comment         string `db:"comment"`
}

package models

type Appointments struct {
	AppointmentId int64  `db:"appointment_id"`
	ClinicId      int64  `db:"clinic_id"`
	PatientId     int64  `db:"patient_id"`
	EmployeeId    int64  `db:"employee_id"`
	StartAt       string `db:"start_at"`
	EndAt         string `db:"end_at"`
	CreatedAt     string `db:"created_at"`
	Status        string `db:"status"`
	Comment       string `db:"comment"`
}
type CheckClinicEmployee struct {
	ID int64 `db:"id"`
}

type DurationAmount struct {
	DurationMinutes int64   `db:"total_duration"`
	TotalPrice      float32 `db:"total_price"`
}

package models

type Clinic struct {
	ClinicID      int64  `db:"clinic_id"`
	ClinicAddress string `db:"clinic_address"`
	Phone         string `db:"phone"`
	Email         string `db:"email"`
	OpeningHours  string `db:"opening_hours"`
	CreatedAt     string `db:"created_at"`
}

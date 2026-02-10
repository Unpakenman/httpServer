package models

import "time"

type Clinic struct {
	ClinicID      int64     `db:"clinic_id"`
	ClinicAddress string    `db:"clinic_adress"`
	Phone         string    `db:"phone"`
	Email         string    `db:"email"`
	OpeningHours  string    `db:"opening_hours"`
	CreatedAt     time.Time `db:"created_at"`
}

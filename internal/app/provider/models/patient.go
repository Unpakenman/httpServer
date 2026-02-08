package models

import "time"

type Patients struct {
	PatientID        int64     `db:"patient_id" `
	FirstName        string    `db:"first_name"`
	LastName         string    `db:"last_name"`
	MiddleName       *string   `db:"middle_name"`
	DocType          int64     `db:"doc_type"`
	DocSeries        int64     `db:"doc_series"`
	DocNumber        int64     `db:"doc_number"`
	Sex              string    `db:"sex"`
	BirthDate        time.Time `db:"birth_date"`
	Phone            string    `db:"phone"`
	Email            string    `db:"email"`
	RegistrationDate time.Time `db:"registration_date"`
}

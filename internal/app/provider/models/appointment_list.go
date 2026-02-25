package models

import "github.com/shopspring/decimal"

type AppointmentList struct {
	PatientName            string          `db:"patient_name"`
	PatientLastName        string          `db:"patient_last_name"`
	AppointmentDTTM        string          `db:"appointment_dttm"`
	AppointmentDescription string          `db:"description"`
	Price                  decimal.Decimal `db:"price"`
	ClinicAdress           string          `db:"clinic_address"`
	DoctorName             string          `db:"doctor_name"`
	DoctorLastName         string          `db:"doctor_last_name"`
}

{{define "CreatePatient"}}
INSERT INTO clinics.patients (
    first_name,
    last_name,
    middle_name,
    doc_type,
    doc_series,
    doc_number,
    sex,
    birth_date,
    phone,
    email,
    registration_date
)VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())
RETURNING {{template "Patients"}}
{{end}}


{{define "CreateClinic"}}
INSERT INTO clinics.clinics(
    clinic_address,
    email,
    opening_hours
    created_at
)VALUES ($1, $2, $3, NOW())
RETURNING {{template "Clinics"}}
{{end}}
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
    opening_hours,
    phone,
    created_at
)VALUES ($1, $2, $3, $4, NOW())
RETURNING {{template "Clinics"}}
{{end}}

{{define "CreateEmployee"}}
INSERT INTO clinics.employees(
    role_id,
    specialization_id,
    first_name,
    last_name,
    middle_name,
    birthdate,
    phone,
    email,
    hire_date
)VALUES($1, $2, $3, $4, $5, $6, $7, $8, NOW())
RETURNING {{template "Employees"}}
{{end}}

{{define "CreateAppointment"}}
INSERT INTO clinics.appointments(
  clinic_id,
  patient_id,
  employee_id,
  appointment_dttm,
  comment
)VALUES($1, $2, $3, $4, $5)
RETURNING {{template "Appointments"}}
{{end}}



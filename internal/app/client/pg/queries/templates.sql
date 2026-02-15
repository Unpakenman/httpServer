{{define "Patients"}}
    patient_id,
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
{{end}}

{{define "Clinics"}}
    clinic_id,
    clinic_address,
    email,
    opening_hours,
    created_at
{{end}}

{{define "Employees"}}
    employee_id,
    role_id,
    specialization_id,
    first_name,
    last_name,
    middle_name,
    birthdate,
    phone,
    email,
    hire_date
 {{end}}

{{define "Appointments"}}
   appointment_id,
   clinic_id,
   patient_id,
   employee_id,
   appointment_dttm,
   created_at,
   status,
   comment
{{end}}


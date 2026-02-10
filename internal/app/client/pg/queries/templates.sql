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
    opening_hours
    created_at
{{end}}
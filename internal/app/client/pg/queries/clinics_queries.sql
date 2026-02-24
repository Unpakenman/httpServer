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


{{define "CreateNewService"}}
INSERT INTO clinics.services(
  name,
  description,
  price,
  duration_minutes,
  is_active
)VALUES($1, $2, $3, $4, $5)
{{end}}

{{define "AppointmentsList"}}
SELECT p.first_name AS patient_name, p.last_name AS patient_last_name,
       a.appointment_dttm, s.description , s.price, c.clinic_address, e.first_name AS doctor_name, e.last_name AS doctor_last_name
FROM clinics.appointments a
         JOIN clinics.patients p on a.patient_id = p.patient_id
         JOIN clinics.clinics c on c.clinic_id = a.clinic_id
         JOIN clinics.appointments_services as2 on as2.appointment_id = a.appointment_id
         JOIN clinics.services s on s.service_id = as2.service_id
         JOIN  clinics.employees e on e.employee_id = a.employee_id
WHERE a.appointment_dttm >= date_trunc('day', now() AT TIME ZONE 'Europe/Moscow')
  AND a.appointment_dttm < date_trunc('day', now() AT TIME ZONE 'Europe/Moscow') + interval '1 day'
ORDER BY a.appointment_dttm
{{end}}



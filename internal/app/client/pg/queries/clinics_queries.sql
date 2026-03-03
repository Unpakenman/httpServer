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
  start_at,
  end_at,
  comment
)VALUES($1, $2, $3, $4, $5, $6)
RETURNING {{template "Appointments"}}
{{end}}

{{define "GetDurationMinutesByServicesIds"}}
SELECT COALESCE(SUM(duration_minutes), 0) as total_duration,
       COALESCE(SUM(price), 0) as total_price
FROM clinics.services
WHERE service_id = ANY($1)
  AND is_active = TRUE;
{{end}}

{{define "CreateTransaction"}}
  INSERT INTO clinics.transactions(
  patient_id,
  clinic_id,
  appointment_id,
  amount,
  discount,
  total_amount,
  services_ids
  )VALUES ($1, $2, $3, $4, $5, $6, $7)
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
       a.start_at, s.description , s.price, c.clinic_address, e.first_name AS doctor_name, e.last_name AS doctor_last_name
FROM clinics.appointments a
         JOIN clinics.patients p on a.patient_id = p.patient_id
         JOIN clinics.clinics c on c.clinic_id = a.clinic_id
         JOIN clinics.appointments_services as2 on as2.appointment_id = a.appointment_id
         JOIN clinics.services s on s.service_id = as2.service_id
         JOIN  clinics.employees e on e.employee_id = a.employee_id
WHERE a.start_at >= date_trunc('day', now() AT TIME ZONE 'Europe/Moscow')
  AND a.start_at < date_trunc('day', now() AT TIME ZONE 'Europe/Moscow') + interval '1 day'
ORDER BY a.start_at
{{end}}

{{define "CheckClinicEmployee"}}
select {{template "ClinicEmployee"}}
from clinics.clinics_employees
where clinic_id = $1 and employee_id=$2
{{end}}


{{define "AppointmentsSlots"}}
  WITH work_time AS (
    SELECT
    tstzrange(
    $2::date + ws.start_working_at,
    $2::date + ws.end_working_at,
    '[)'
    ) AS work_range,
    tstzrange(
    $2::date + ws.break_start_at,
    $2::date + ws.break_end_at,
    '[)'
    ) AS break_range
    FROM clinics.work_schedule ws
    WHERE ws.employee_id = $1
    AND EXTRACT(DOW FROM $2::date) = ws.day_of_week
    ),
    busy_time AS (
    SELECT range_agg(tstzrange(start_at, end_at, '[)')) AS busy_ranges
    FROM clinics.appointments
    WHERE employee_id = $1
    AND start_at::date = $2
    AND status = 'confirmed'
    ),
    exceptions AS (
    SELECT range_agg(tstzrange(exception_start_at, exception_end_at, '[)')) AS ex_ranges
    FROM clinics.working_exceptions
    WHERE employee_id = $1
    AND exception_start_at::date = $2
    ),
    free AS (
    SELECT
    (
    multirange(work_range)
    - COALESCE(busy_ranges, '{}'::tstzmultirange)
    - COALESCE(ex_ranges, '{}'::tstzmultirange)
    - multirange(break_range)
    ) AS free_ranges
    FROM work_time, busy_time, exceptions
    ),
    expanded AS (
    SELECT unnest(free_ranges) AS r
    FROM free
    ),
    hour_slots AS (
    SELECT
    generate_series(
    lower(r),
    upper(r) - interval '1 hour',
    interval '1 hour'
    ) AS slot_start
    FROM expanded
    )
SELECT
    slot_start,
    slot_start + interval '1 hour' AS slot_end
FROM hour_slots
ORDER BY slot_start;
{{end}}

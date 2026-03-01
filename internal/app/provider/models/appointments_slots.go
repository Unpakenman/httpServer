package models

type AppointmentsSlots struct {
	StartAt string `db:"slot_start"`
	EndAt   string `db:"slot_end"`
}

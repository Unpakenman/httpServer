package clinics

type AddClinicRequest struct {
	ClinicAdress string
	Phone        string
	Email        string
	OpeningHours string
}

type AddClinicResponse struct {
	ClinicId int64
}

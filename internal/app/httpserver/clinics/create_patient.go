package clinics

import (
	"encoding/json"
	"httpServer/internal/app/httpserver/models"
	"io"
	"net/http"
)

func (r *httpRouter) CreatePatient(w http.ResponseWriter, req *http.Request) {
	bodyBytes, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		r.logger.ErrorCtx(req.Context(), readErr, "failed to read request body: %w", readErr)
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}
	var request models.CreatePatientRequest
	requestErr := json.Unmarshal(bodyBytes, &request)
	if requestErr != nil {
		r.logger.ErrorCtx(req.Context(), requestErr, "failed to unmarshal request: %w", requestErr)
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	if validateErrors := r.validator.CreatePatient(request); validateErrors != nil {
		r.logger.ErrorCtx(req.Context(), validateErrors, "failed to validate patient: %w", validateErrors)
		return
	}
	requestCreatePatient := r.mapper.HttpToCreatePayinRequest(request)
	w.Header().Set("Content-Type", "application/json")
	response, err := r.usecase.CreatePatient(req.Context(), requestCreatePatient)
	if err != nil {
		r.logger.ErrorCtx(req.Context(), err, "failed to create patient:", err)
		w.WriteHeader(http.StatusBadRequest)
		if encodeErr := json.NewEncoder(w).Encode(models.ErrorResponse{
			Error: err.Error(),
		}); encodeErr != nil {
			r.logger.ErrorCtx(req.Context(), encodeErr, "failed to encode response:", encodeErr)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	patient := r.mapper.CreatePatientToHttp(response)

	if encodeErr := json.NewEncoder(w).Encode(models.CreatePatientResponse{
		Status:    "SUCCESS",
		PatientID: patient.PatientID,
	}); encodeErr != nil {
		r.logger.ErrorCtx(req.Context(), encodeErr, "failed to encode response:", encodeErr)
	}
}

package clinics

import (
	"encoding/json"
	"httpServer/internal/app/httpserver/models"
	"io"
	"net/http"
)

func (r *httpRouter) CreatePatient(w http.ResponseWriter, req *http.Request) {
	bodyBytes, readErr := io.ReadAll(req.Body)
	r.logger.Info("CreatePatient", bodyBytes)
	if readErr != nil {
		r.logger.ErrorContext(req.Context(), "failed to read request body: %w", readErr)
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}
	var request models.CreatePatientRequest
	requestErr := json.Unmarshal(bodyBytes, &request)
	if requestErr != nil {
		r.logger.ErrorContext(req.Context(), "failed to unmarshal request: %w", requestErr)
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	if validateErrors := r.validator.CreatePatient(request); validateErrors != nil {
		r.logger.ErrorContext(req.Context(), "failed to validate patient: %w", validateErrors)
		return
	}
	requestCreatePatient := r.mapper.HttpToCreatePayinRequest(request)
	w.Header().Set("Content-Type", "application/json")
	response, err := r.usecase.CreatePatient(req.Context(), requestCreatePatient)
	if err != nil {
		r.logger.ErrorContext(req.Context(), "failed to create patient: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		if encodeErr := json.NewEncoder(w).Encode(err.Error());
		//if encodeErr := json.NewEncoder(w).Encode(models.CreatePatientResponse{
		//Status: "FAILED",	}
		encodeErr != nil {
			r.logger.ErrorContext(req.Context(), "failed to encode response: %w", encodeErr)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	patientId := r.mapper.CreatePatientToHttp(response)
	if encodeErr := json.NewEncoder(w).Encode(models.CreatePatientResponse{
		Status:    "SUCCESS",
		PatientId: patientId.PatientId,
	}); encodeErr != nil {
		r.logger.ErrorContext(req.Context(), "failed to encode response: %w", encodeErr)
	}
}

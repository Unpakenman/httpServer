package clinics

import (
	"encoding/json"
	"fmt"
	"httpServer/internal/app/httpserver/models"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func (r *httpRouter) CreatePatient(w http.ResponseWriter, req *http.Request) {
	bodyBytes, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		r.logger.ErrorCtx(req.Context(), readErr, "failed to read request body")
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}
	var request models.CreatePatientRequest
	requestErr := json.Unmarshal(bodyBytes, &request)
	if requestErr != nil {
		r.logger.ErrorCtx(req.Context(), requestErr, "failed to unmarshal request")
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	if err := r.validator.CreatePatient(request); err != nil {
		r.logger.ErrorCtx(req.Context(), fmt.Errorf("failed in validate body: %w", err))
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}
	var phoneRegex = regexp.MustCompile(`^(?:\+7|8)?9\d{2}\d{7}$`)

	if !phoneRegex.MatchString(request.PhoneNumber) {
		r.logger.ErrorCtx(req.Context(), fmt.Errorf("invalid phone number: %s", request.PhoneNumber))
		http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
		return
	}
	requestCreatePatient := r.mapper.HttpToCreatePayinRequest(request)
	w.Header().Set("Content-Type", "application/json")
	response, err := r.usecase.CreatePatient(req.Context(), requestCreatePatient)
	if err != nil {
		r.logger.ErrorCtx(req.Context(), err, "failed to create patient")
		w.WriteHeader(http.StatusInternalServerError)
		if encodeErr := json.NewEncoder(w).Encode(models.ErrorResponse{
			Error: err.Error(),
		}); encodeErr != nil {
			r.logger.ErrorCtx(req.Context(), encodeErr, "failed to encode response")
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	patientID := strconv.FormatInt(response.PatientID, 10)

	if encodeErr := json.NewEncoder(w).Encode(models.CreatePatientResponse{
		PatientID: &patientID,
	}); encodeErr != nil {
		r.logger.ErrorCtx(req.Context(), encodeErr, "failed to encode response")
	}
}

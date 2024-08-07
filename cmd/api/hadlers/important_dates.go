package handlers

import (
	"net/http"
	"time"

	"github.com/caiquetorres/amparo/cmd/api/dtos"
	"github.com/caiquetorres/amparo/cmd/api/validators"
	"github.com/caiquetorres/amparo/internal"
)

type ImportantDatesHandler struct{}

func NewImportantDatesHandler() *ImportantDatesHandler {
	return &ImportantDatesHandler{}
}

func (h *ImportantDatesHandler) HandleImportantDatesPost(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	dateOfDeath, err := validators.ParseDate(dateStr)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		internal.WriteJSON(w, http.StatusBadRequest, internal.NewApiError("Invalid date"))
		return
	}
	if dateOfDeath.After(time.Now()) {
		w.Header().Add("Content-Type", "application/json")
		internal.WriteJSON(w, http.StatusUnprocessableEntity, internal.NewApiError("The date cannot be in the future"))
		return
	}
	scheduleMass := dateOfDeath.AddDate(0, 0, 3)
	registerDeath := dateOfDeath.AddDate(0, 0, 15)
	pensionRequest := dateOfDeath.AddDate(0, 0, 90)
	insuranceClaim := dateOfDeath.AddDate(0, 0, 365)
	w.Header().Add("Content-Type", "application/json")
	internal.WriteJSON(w, http.StatusOK, dtos.ImportantDatesResponse{
		ScheduleMass:   scheduleMass.Format("2006-01-02"),
		RegisterDeath:  registerDeath.Format("2006-01-02"),
		PensionRequest: pensionRequest.Format("2006-01-02"),
		InsuranceClaim: insuranceClaim.Format("2006-01-02"),
	})
}

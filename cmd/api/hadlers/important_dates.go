package handlers

import (
	"encoding/json"
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
	var importantDates dtos.GetImportantDates
	err := json.NewDecoder(r.Body).Decode(&importantDates)

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		internal.WriteJSON(w, http.StatusBadRequest, internal.NewApiError("Invalid JSON body"))
		return
	}

	dateOfDeathStr := importantDates.DateOfDeath

	if dateOfDeathStr == "" {
		internal.WriteJSON(w, http.StatusBadRequest, internal.NewApiError("Missing date_of_death property"))
		return
	}

	dateOfDeath, err := validators.ParseDate(dateOfDeathStr)

	if err != nil {
		internal.WriteJSON(w, http.StatusBadRequest, internal.NewApiError("Invalid date_of_death property"))
		return
	}

	if dateOfDeath.After(time.Now()) {
		internal.WriteJSON(w, http.StatusUnprocessableEntity, internal.NewApiError("The date_of_death cannot be in the future"))
		return
	}

	scheduleMass := dateOfDeath.AddDate(0, 0, 3)
	registerDeath := dateOfDeath.AddDate(0, 0, 15)
	pensionRequest := dateOfDeath.AddDate(0, 0, 90)
	insuranceClaim := dateOfDeath.AddDate(0, 0, 365)

	internal.WriteJSON(w, http.StatusOK, dtos.ImportantDatesResponse{
		ScheduleMass:   scheduleMass.Format("2006-01-02"),
		RegisterDeath:  registerDeath.Format("2006-01-02"),
		PensionRequest: pensionRequest.Format("2006-01-02"),
		InsuranceClaim: insuranceClaim.Format("2006-01-02"),
	})
}

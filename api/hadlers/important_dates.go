package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/caiquetorres/amparo/api/dtos"
	"github.com/caiquetorres/amparo/api/validators"
	"github.com/caiquetorres/amparo/internal"
)

type ImportantDatesHandler struct{}

func NewImportantDatesHandler() *ImportantDatesHandler {
	return &ImportantDatesHandler{}
}

func (h *ImportantDatesHandler) HandleImportantDatesPost(w http.ResponseWriter, r *http.Request) {
	// Get the payload
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

	dates := []dtos.ImportantDateResponse{
		{
			Name: "schedule_mass",
			Date: dateOfDeath.AddDate(0, 0, 3).Format("2006-01-02"),
		},
		{
			Name: "register_deach",
			Date: dateOfDeath.AddDate(0, 0, 15).Format("2006-01-02"),
		},
		{
			Name: "pension_request",
			Date: dateOfDeath.AddDate(0, 0, 90).Format("2006-01-02"),
		},
		{
			Name: "insurange_claim",
			Date: dateOfDeath.AddDate(0, 0, 365).Format("2006-01-02"),
		},
	}

	internal.WriteJSON(w, http.StatusOK, dates)
}

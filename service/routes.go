package service

import (
	"log"
	"net/http"
	"time"

	"github.com/caiquetorres/amparo/internal"
	"github.com/gorilla/mux"
)

type Handler struct{}

type ImportantDates struct {
	Date1 string `json:"date1"`
	Date2 string `json:"date2"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/important-dates", h.handleImportantDatesRoute).Methods("POST")
	log.Println("Mapped POST /important-dates")
}

func (h *Handler) handleImportantDatesRoute(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := parseDate(dateStr)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		internal.WriteJSON(w, http.StatusBadRequest, internal.NewApiError("Invalid date"))
		return
	}
	if date.After(time.Now()) {
		w.Header().Add("Content-Type", "application/json")
		internal.WriteJSON(w, http.StatusUnprocessableEntity, internal.NewApiError("The date cannot be in the future"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	internal.WriteJSON(w, http.StatusOK, ImportantDates{
		Date1: "date1",
		Date2: "date2",
	})
}

func parseDate(date string) (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, date)
}

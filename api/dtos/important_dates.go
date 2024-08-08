package dtos

type GetImportantDates struct {
	DateOfDeath string `json:"date_of_death"`
}

type ImportantDateResponse struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

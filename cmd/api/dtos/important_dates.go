package dtos

type ImportantDatesResponse struct {
	ScheduleMass   string `json:"schedule_mass"`
	RegisterDeath  string `json:"register_death"`
	PensionRequest string `json:"pension_request"`
	InsuranceClaim string `json:"insurance_claim"`
}
package dto

type EventCertificationRequestDto struct {
	Certified           bool `json:"certified,omitempty"`
	ToggleCertification bool `json:"toggle_certification,omitempty"`
}

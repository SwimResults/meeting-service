package dto

import "time"

type IncidentDateRequest struct {
	Time           time.Time `json:"time"`
	UpdateTimeZone bool      `json:"update_time_zone"`
}

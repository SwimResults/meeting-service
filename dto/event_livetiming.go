package dto

import "github.com/swimresults/meeting-service/model"

type EventLivetimingDto struct {
	Event     model.Event `json:"event,omitempty"`
	PrevEvent model.Event `json:"prev_event,omitempty"`
	NextEvent model.Event `json:"next_event,omitempty"`
}

package dto

import "github.com/swimresults/meeting-service/model"

type ImportEventRequestDto struct {
	Event             model.Event `json:"event"` // number, gender, meeting, (relay-)distance, ordering
	StyleName         string      `json:"style_name"`
	MeetingPartNumber int         `json:"part_number"`
}

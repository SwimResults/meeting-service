package dto

import "github.com/swimresults/meeting-service/model"

type ImportAgeGroupRequestDto struct {
	AgeGroup model.AgeGroup `json:"age_group"`
}

package dto

import "github.com/swimresults/meeting-service/model"

// ImportAgeGroupRequestDto is deprecated: moved to start service as ranking
type ImportAgeGroupRequestDto struct {
	AgeGroup model.AgeGroup `json:"age_group"`
}

package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/swimresults/meeting-service/model"
	"testing"
)

func TestSetAgesForAgeGroup(t *testing.T) {
	group := model.AgeGroup{
		MinAge: "2004",
		MaxAge: "2002",
		Ages:   nil,
		IsYear: true,
	}

	SetAgesForAgeGroup(&group)

	assert.Equal(t, []int{2002, 2003, 2004}, group.Ages)
}

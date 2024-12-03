package client

import (
	"encoding/json"
	"fmt"
	"github.com/swimresults/meeting-service/dto"
	"github.com/swimresults/meeting-service/model"
	"github.com/swimresults/service-core/client"
	"net/http"
	"strconv"
)

type AgeGroupClient struct {
	apiUrl string
}

func NewAgeGroupClient(url string) *AgeGroupClient {
	return &AgeGroupClient{apiUrl: url}
}

func (c *AgeGroupClient) ImportAgeGroup(ageGroup model.AgeGroup) (*model.AgeGroup, bool, error) {
	request := dto.ImportAgeGroupRequestDto{
		AgeGroup: ageGroup,
	}

	res, err := client.Post(c.apiUrl, "age_group/import", request, nil)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

	newAgeGroup := &model.AgeGroup{}
	err = json.NewDecoder(res.Body).Decode(newAgeGroup)
	if err != nil {
		return nil, false, err
	}

	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("import age group request returned: %d", res.StatusCode)
	}
	return newAgeGroup, res.StatusCode == http.StatusCreated, nil
}

func (c *AgeGroupClient) GetAgeGroupsForMeetingAndEvent(meeting string, number int) (*[]model.AgeGroup, error) {
	fmt.Printf("request '%s'\n", c.apiUrl+"/age_group/meet/"+meeting+"/event/"+strconv.Itoa(number))

	res, err := client.Get(c.apiUrl, "age_group/meet/"+meeting+"/event/"+strconv.Itoa(number), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetAgeGroupsForMeetingAndEvent received error: %d\n", res.StatusCode)
	}

	ageGroups := &[]model.AgeGroup{}
	err = json.NewDecoder(res.Body).Decode(ageGroups)
	if err != nil {
		return nil, err
	}

	return ageGroups, nil
}

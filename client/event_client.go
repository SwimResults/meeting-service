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

type EventClient struct {
	apiUrl string
}

func NewEventClient(url string) *EventClient {
	return &EventClient{apiUrl: url}
}

func (c *EventClient) ImportEvent(event model.Event, styleName string, PartNumber int) (*model.Event, bool, error) {
	request := dto.ImportEventRequestDto{
		Event:             event,
		StyleName:         styleName,
		MeetingPartNumber: PartNumber,
	}

	res, err := client.Post(c.apiUrl, "event/import", request, nil)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

	newEvent := &model.Event{}
	err = json.NewDecoder(res.Body).Decode(newEvent)
	if err != nil {
		return nil, false, err
	}

	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("import event request returned: %d", res.StatusCode)
	}
	return newEvent, res.StatusCode == http.StatusCreated, nil
}

func (c *EventClient) GetEventByMeetingAndNumber(meeting string, number int) (*model.Event, error) {
	fmt.Printf("request '%s'\n", c.apiUrl+"event/meet/"+meeting+"/event/"+strconv.Itoa(number))

	res, err := client.Get(c.apiUrl, "event/meet/"+meeting+"/event/"+strconv.Itoa(number), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetAthleteByNameAndYear received error: %d\n", res.StatusCode)
	}

	event := &model.Event{}
	err = json.NewDecoder(res.Body).Decode(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

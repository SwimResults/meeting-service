package client

import (
	"encoding/json"
	"fmt"
	"github.com/swimresults/meeting-service/dto"
	"github.com/swimresults/meeting-service/model"
	"github.com/swimresults/service-core/client"
	"net/http"
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

	res, err := client.Post(c.apiUrl, "event/import", request)
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

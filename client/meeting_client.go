package client

import (
	"encoding/json"
	"fmt"
	"github.com/swimresults/meeting-service/model"
	"github.com/swimresults/service-core/client"
	"net/http"
)

type MeetingClient struct {
	apiUrl string
}

func NewMeetingClient(url string) *MeetingClient {
	return &MeetingClient{apiUrl: url}
}

func (c *MeetingClient) GetMeetingById(id string) (*model.Meeting, error) {
	fmt.Printf("request '%s'\n", c.apiUrl+"meeting/meet/"+id)

	res, err := client.Get(c.apiUrl, "meeting/meet/"+id, nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetMeetingById received error: %d\n", res.StatusCode)
	}

	meeting := &model.Meeting{}
	err = json.NewDecoder(res.Body).Decode(meeting)
	if err != nil {
		return nil, err
	}

	return meeting, nil
}

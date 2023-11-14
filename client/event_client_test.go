package client

//import (
//	"fmt"
//	"github.com/swimresults/meeting-service/model"
//	"testing"
//)
//
//func TestEventClient_ImportEvent(t *testing.T) {
//	client := NewEventClient("http://localhost:8089/")
//
//	event := model.Event{
//		Number:   4,
//		Distance: 100,
//		Meeting:  "IESC19",
//		Gender:   "m",
//	}
//
//	r, _, e := client.ImportEvent(event, "Schmetterling", 1)
//	if e != nil {
//		fmt.Printf(e.Error())
//	}
//	fmt.Println(r)
//}

//func TestEventClient_GetEventByMeetingAndNumber(t *testing.T) {
//	client := NewEventClient("https://api.swimresults.de/meeting/v1/")
//
//	r, e := client.GetEventByMeetingAndNumber("IESC22", 13)
//	if e != nil {
//		fmt.Printf(e.Error())
//	}
//	fmt.Println(r)
//}

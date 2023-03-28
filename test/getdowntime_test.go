package api

import (
	"datadog/api"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetdowntime(t *testing.T) {

	out := api.DowntimeStruct{}
	jsonFile, err := ioutil.ReadFile("downtimes.json")
	if err != nil {
		t.Logf("Unable to Call function: %+v", err)
	}
	// Convert Json to Struct
	err = json.Unmarshal(jsonFile, &out)
	if err != nil {
		t.Errorf("Unable to Unmarshal json file %+v", err)
	}

	expectedID := 2806375
	expectedActive := true
	expectedDisabled := false
	expectedMessage := "All alerting will be silenced due to a Deployment @erastus.thambo@broadridge.com"
	expectedStart := 1679511431
	expectedEnd := 1679515031

	for _, downtime := range out {
		if *downtime.CreatorID != expectedID {
			t.Fatalf(" active %d. Got %d", expectedID, *downtime.CreatorID)
		}
		if *downtime.Active != expectedActive {
			t.Fatalf(" active %t. Got %v", expectedActive, *downtime.Active)
		}
		if *downtime.Disabled != expectedDisabled {
			t.Fatalf(" active %t. Got %v", expectedDisabled, *downtime.Disabled)
		}
		if *downtime.Message != expectedMessage {
			t.Fatalf(" active %s. Got %s", expectedMessage, *downtime.Message)
		}
		if *downtime.Start != expectedStart {
			t.Fatalf(" active %d. Got %d", expectedStart, *downtime.Start)
		}
		if *downtime.End != expectedEnd {
			t.Fatalf(" active %d. Got %d", expectedEnd, *downtime.End)
		}

	}

}

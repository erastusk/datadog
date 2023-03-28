package api

import (
	"context"
	"datadog/env"
	"encoding/json"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func Getdowntime(s *env.EnvVariablesLoad) ([]byte, error) {
	// Get all downtimes returns "OK" response
	s.Logger.Printf("Getting a Listing of All Downtimes\n")
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDowntimesApi(apiClient)
	body, r, err := api.ListDowntimes(ctx, *datadogV1.NewListDowntimesOptionalParameters())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	responseContent, err := json.Marshal(body)
	if err != nil {
		s.Logger.Printf("Could not Marshal downtimes but will continue with downtime creation\n")
		return nil, err
	}
	// _ = ioutil.WriteFile("downtimes.json", responseContent, 0644)
	return responseContent, nil

}

func GetDowntimeIds(s *env.EnvVariablesLoad, body []byte) []int {
	s.Logger.Printf("Retrieving Downtime IDs based on creator id: %+v\n", s.CreatorID)
	t, err := strconv.Atoi(s.CreatorID)
	if err != nil {
		s.Logger.Printf("%v\n", err)
		t = 2797870
	}
	out := DowntimeStruct{}
	var downtime_ids []int
	err = json.Unmarshal(body, &out)

	if err != nil {
		s.Logger.Printf("%+v\n", err)
	}
	for _, downtime := range out {
		if len(downtime.MonitorTags) > 2 {
			if t == *downtime.CreatorID && *downtime.Active && downtime.MonitorTags[0] == s.Tags[2] {
				downtime_ids = append(downtime_ids, *downtime.Id)
			}
		}
	}
	return downtime_ids

}

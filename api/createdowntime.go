package api

import (
	"context"
	"datadog/env"
	"encoding/json"
	"strconv"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func Downtime(s *env.EnvVariablesLoad) ([]byte, int, error) {
	t, err := strconv.Atoi(s.Time)
	if err != nil {
		s.Logger.Printf("%v\n", err)
		t = 3
	}

	body := datadogV1.Downtime{
		Message: &s.Message,
		Scope: []string{
			"*",
		},
		Start:                         datadog.PtrInt64(time.Now().Unix()),
		End:                           *datadog.NewNullableInt64(datadog.PtrInt64(time.Now().Add(time.Hour * time.Duration(t)).Unix())),
		Timezone:                      datadog.PtrString("EST"),
		MuteFirstRecoveryNotification: datadog.PtrBool(true),
		MonitorTags:                   s.Tags,
	}

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDowntimesApi(apiClient)
	resp, r, err := api.CreateDowntime(ctx, body)

	if err != nil {
		s.Logger.Printf("Full HTTP response: %v\n", r)
		return nil, r.StatusCode, err
	}
	defer r.Body.Close()
	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	return responseContent, r.StatusCode, nil
}

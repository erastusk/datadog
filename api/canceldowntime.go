package api

import (
	"context"
	"datadog/env"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func Canceltdowntime(s *env.EnvVariablesLoad, downtimeID int64) ([]byte, error) {

	DowntimeID := downtimeID
	s.Logger.Printf("Deleting Downtime id: %+v...\n", downtimeID)

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDowntimesApi(apiClient)
	r, err := api.CancelDowntime(ctx, DowntimeID)

	if err != nil {
		s.Logger.Printf("Full HTTP response: %v\n", r)
		return nil, err
	}
	defer r.Body.Close()
	return nil, nil

}

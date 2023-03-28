package api

import (
	"context"
	"datadog/env"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func ValidateKeys(s *env.EnvVariablesLoad) int {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAuthenticationApi(apiClient)
	_, r, err := api.Validate(ctx)
	if err != nil {
		s.Logger.Printf("Full HTTP response: %v\n", r.StatusCode)
	}
	defer r.Body.Close()
	if r.StatusCode == 403 {
		return r.StatusCode
	}
	if r.StatusCode == 429 || r.StatusCode >= 500 {
		sleep := 1
		for i := 0; i < 6; i++ {
			if i > 0 {
				s.Logger.Printf("retrying API Key validation ......")
				time.Sleep(time.Duration(sleep) * time.Second)
				sleep *= 2
			}
			_, r, err := api.Validate(ctx)
			if err != nil {
				s.Logger.Printf("Failure while retrying: %v\n", r)
			}
			if r.StatusCode == 200 {
				return r.StatusCode
			}
		}
	}
	return r.StatusCode
}

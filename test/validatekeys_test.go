package api

import (
	"datadog/api"
	"datadog/env"
	"log"
	"os"
	"testing"
)

var logger = log.New(os.Stdout, "INFO... ", log.Ldate|log.Ltime|log.Lshortfile)

// var err = env.LoadVars()
var s = &env.EnvVariablesLoad{
	Tags: []string{
		"env:dev",
		"app:workstationframework",
	},
	CreatorID: "2806375",
	Message:   "All alerting will be silenced due to a Deployment @erastus.thambo@broadridge.com",
	Logger:    logger,
	Time:      "1",
}

func TestValidateKeys(t *testing.T) {
	// if err != nil {
	// 	t.Logf("App Load environment variables failed: %+v", err)
	// }

	result := api.ValidateKeys(s)
	if result != 200 {
		t.Errorf("ValidateKeys FAILED. Expected response code 200, got %d", result)
	} else {
		t.Logf("ValidateKeys PASSED. Expected response code 200, got %d", result)
	}
}

package main

import (
	"datadog/api"
	"datadog/env"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(request events.CloudWatchEvent) {
	logger := log.New(os.Stdout, "INFO... ", log.Ldate|log.Ltime|log.Lshortfile)
	s := env.NewEnvLoad(logger)

	// Before proceeding, ensure DD_APP_KEY has been passed as an environment variable
	s.Logger.Printf("**** Starting Datadog Downtime Creation ****\n\n")
	if os.Getenv("DD_APP_KEY") == "" {
		s.Logger.Fatalf("Could not find environment variable DD_APP_KEY, exiting....\n")
	}

	// Check with Datadog if APP Key is valid
	s.Logger.Printf("Validating environment variable (DD_APP_KEY) APP key\n")
	valid := api.ValidateKeys(s)
	if valid != 200 {
		s.Logger.Fatalf("Unable to validate API Key.....\n")
	}
	s.Logger.Printf("API Key is Valid\n")

	//Retrieve All Active Downtimes in Datadog
	s.Logger.Printf("Checking for existing downtimes...\n")
	resp, err := api.Getdowntime(s)
	if err != nil {
		s.Logger.Printf("%+v\n", err)
	}

	//Each APP KEY has a unique creator ID, check if there's an active Downtime created with our ID
	delIds := api.GetDowntimeIds(s, resp)
	if len(delIds) > 0 {
		s.Logger.Printf("Downtime(s) found with creator id :%s -> %+v Exiting...\n**** End Datadog Downtime Creation ****\n", s.CreatorID, delIds)
		return
	}
	s.Logger.Printf("Creating Downtime with tags --->%s\n", s.Tags)
	respCr, status, err := api.Downtime(s)
	if err != nil {
		s.Logger.Printf("%+v\n", err)
	}
	s.Logger.Printf("Successfully created Downtime...\n%s\n Response code: %d", string(respCr), status)
	s.Logger.Printf("**** End Datadog Downtime Creation ****\n")

}

package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariablesLoad struct {
	EnvFile   string
	Tags      []string
	CreatorID string
	Message   string
	Logger    *log.Logger
	Time      string
}

func NewEnvLoad(logger *log.Logger) *EnvVariablesLoad {
	mess := "All alerting will be silenced due to a Deployment "
	return &EnvVariablesLoad{
		Tags: []string{
			"env:" + os.Getenv("env"),
			"app:" + os.Getenv("app"),
			"alias:" + os.Getenv("alias"),
		},
		Message:   mess + os.Getenv("message"),
		Logger:    logger,
		CreatorID: os.Getenv("creator_id"),
		Time:      os.Getenv("time"),
	}
}

func TestNewEnvLoad(logger *log.Logger) *EnvVariablesLoad {
	return &EnvVariablesLoad{
		Tags: []string{
			"env:dev",
			"app:workstationframework",
			"alias:dev2",
		},
		CreatorID: "2806375",
		Message:   "All alerting will be silenced due to a Deployment @erastus.thambo@broadridge.com",
		Logger:    logger,
		Time:      "1",
	}
}

func LoadVars() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil

}

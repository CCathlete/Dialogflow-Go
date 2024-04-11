package dialogflowUtils

import (
	"context"
	"log"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"google.golang.org/api/option"
)

/*
	DialogflowProcessor has all the information for connecting with

google's DialogFlow.
*/
type DialogflowProcessor struct {
	projectID        string
	authJSONFilePath string
	lang             string
	timeZone         string
	sessionClient    *dialogflow.SessionsClient
	ctx              context.Context
}

// NLPResponse is the struct for the response.
type NLPResponse struct {
	Intent     string            `json:"intent"`
	Confidence float32           `json:"confidence"`
	Entities   map[string]string `json:"entities"`
}

var Dfp DialogflowProcessor

func (dfp *DialogflowProcessor) init(a ...string) (err error) {
	dfp.projectID = a[0]
	dfp.authJSONFilePath = a[1]
	dfp.lang = a[2]
	dfp.timeZone = a[3]

	// Auth process: https://dialogflow.com/docs/reference/v2-auth-setup

	dfp.ctx = context.Background()
	sessionClient, err := dialogflow.NewSessionsClient(dfp.ctx,
		option.WithCredentialsFile(dfp.authJSONFilePath))
	if err != nil {
		log.Fatal("Error in auth with Dialogflow.")
	}
	dfp.sessionClient = sessionClient
	return err
}

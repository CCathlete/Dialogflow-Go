package dialogflowUtils

import (
	"context"
	"fmt"
	"log"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
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

func (dfp *DialogflowProcessor) Init(a ...string) (err error) {
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

func (dfp *DialogflowProcessor) ProcessNLP(rawMessage string, username string,
) (r NLPResponse) {
	sessionID := username
	request := dialogflowpb.DetectIntentRequest{
		Session: fmt.Sprintf("projects/%s/agent/sessions/%s", dfp.projectID,
			sessionID),
		QueryInput: &dialogflowpb.QueryInput{
			Input: &dialogflowpb.QueryInput_Text{
				Text: &dialogflowpb.TextInput{
					Text:         rawMessage,
					LanguageCode: dfp.lang,
				},
			},
		},
		QueryParams: &dialogflowpb.QueryParameters{
			TimeZone: dfp.timeZone,
		},
	}
	response, err := dfp.sessionClient.DetectIntent(dfp.ctx, &request)
	if err != nil {
		log.Fatalf("Error in communication with Dialogflow %s", err.Error())
		return r
	}
	queryResult := response.GetQueryResult()
	if queryResult.Intent != nil {
		r.Intent = queryResult.Intent.DisplayName
		r.Confidence = float32(queryResult.IntentDetectionConfidence)
	}
	r.Entities = make(map[string]string)
	params := queryResult.Parameters.GetFields()
	if len(params) > 0 {
		for paramName, p := range params {
			fmt.Printf("Param %s: %s (%s)", paramName, p.GetStringValue(),
				p.String())
			extractedValue := extractDialogflowEntities(p)
			r.Entities[paramName] = extractedValue
		}
	}
	return r
}

package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func GetConfig() *aws.Config {
	config := aws.NewConfig()
	config.Region = aws.String(endpoints.UsEast1RegionID)
	return config
}

func GetSession() *session.Session {
	return session.Must(session.NewSession(GetConfig()))
}

func LogAwsError(err error) {
	switch x := err.(type) {
	case awserr.Error:
		log.Printf("AWS error (%s): '%s'", x.Code(), x.Message())
		break
	default:
		log.Printf("Error: '%s'", err.Error())
	}
}

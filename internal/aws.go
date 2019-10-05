package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func GetConfig() *aws.Config {
	return &aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	}
}

func GetSession() *session.Session {
	sess, sessErr := session.NewSession(GetConfig())
	if sessErr != nil {
		log.Panicf("Session error: %s", sessErr.Error())
	}
	return sess
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

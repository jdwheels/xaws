package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func StartInstance(id string, dryRun bool) {
	computeService := ec2.New(GetSession())

	input := ec2.StartInstancesInput{
		DryRun:      aws.Bool(dryRun),
		InstanceIds: []*string{aws.String(id)},
	}

	req, res := computeService.StartInstancesRequest(&input)

	if err := req.Send(); err != nil {
		LogAwsError(err)
		return
	}

	for _, instance := range res.StartingInstances {
		log.Printf("CurrentState: %s", *instance.CurrentState)
		log.Printf("PreviousState: %s", *instance.PreviousState)
	}
}

func StopInstance(id string, dryRun bool) {
	computeService := ec2.New(GetSession())

	input := ec2.StopInstancesInput{
		DryRun:      aws.Bool(dryRun),
		InstanceIds: []*string{aws.String(id)},
	}

	req, res := computeService.StopInstancesRequest(&input)

	if err := req.Send(); err != nil {
		LogAwsError(err)
		return
	}

	for _, instance := range res.StoppingInstances {
		log.Printf("CurrentState: %s", *instance.CurrentState)
		log.Printf("PreviousState: %s", *instance.PreviousState)
	}
}

func InstanceStatus(id string) (statuses []*ec2.InstanceStatus, err error) {
	computeService := ec2.New(GetSession())

	input := ec2.DescribeInstanceStatusInput{
		InstanceIds: []*string{aws.String(id)},
	}

	req, res := computeService.DescribeInstanceStatusRequest(&input)

	if err := req.Send(); err != nil {
		LogAwsError(err)
	} else {
		statuses = res.InstanceStatuses
	}

	return
}

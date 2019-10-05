package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"log"
	"xaws/internal"
)

func CheckIt(name string) (err error) {
	autoscaleService := autoscaling.New(internal.GetSession())

	input := autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{aws.String(name)},
	}

	log.Printf(`Waiting for auto scaling group "%s" to be in service...`, name)
	if err := autoscaleService.WaitUntilGroupInService(&input); err != nil {
		internal.LogAwsError(err)
	}

	return
}

func StartEC2Cluster(asgName string) bool {
	if _, err := internal.AutoScaleIt(asgName, 1); err != nil {
		return false
	} else {
		return CheckIt(asgName) == nil
	}
}

func StopEC2Cluster(asgName string) bool {
	if _, err := internal.AutoScaleIt(asgName, 0); err != nil {
		return false
	} else {
		return true
	}
}

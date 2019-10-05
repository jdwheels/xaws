package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"log"
)

func AutoScaleIt(name string, desiredCapacity int64) (result *autoscaling.UpdateAutoScalingGroupOutput, err error) {
	autoscaleService := autoscaling.New(GetSession())

	input := autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String(name),
		DesiredCapacity:      aws.Int64(desiredCapacity),
	}

	log.Printf(`Bringing auto scaling group "%s" into service...`, name)
	req, res := autoscaleService.UpdateAutoScalingGroupRequest(&input)

	if err := req.Send(); err != nil {
		LogAwsError(err)
	} else {
		log.Printf(`Auto scaling group "%s" scaling instances to "%d"...`, name, desiredCapacity)
		result = res
	}

	return
}

func CheckIt(name string) (err error) {
	autoscaleService := autoscaling.New(GetSession())

	input := autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{aws.String(name)},
	}

	log.Printf(`Waiting for auto scaling group "%s" to be in service...`, name)
	if err := autoscaleService.WaitUntilGroupInService(&input); err != nil {
		LogAwsError(err)
	}

	return
}

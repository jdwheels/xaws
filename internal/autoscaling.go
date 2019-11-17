package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
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
		return nil, err
	} else {
		log.Printf(`Auto scaling group "%s" scaling instances to "%d"...`, name, desiredCapacity)
		result = res
	}

	return result, err
}

func CheckIt(name string) (count int, status string, err error) {
	autoscaleService := autoscaling.New(GetSession())

	input := autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{aws.String(name)},
	}

	//log.Printf(`Waiting for auto scaling group "%s" to be in service...`, name)
	res, err := autoscaleService.DescribeAutoScalingGroups(&input)
	if err != nil {
		LogAwsError(err)
	}

	for _, x := range res.AutoScalingGroups {
		count = len(x.Instances)
		if count > 0 {
			status = *x.Instances[0].LifecycleState
			log.Printf("%+v", *x.Instances[0])
		} else if *x.DesiredCapacity > 0 {
			status = ec2.InstanceStateNamePending
		}
		break
	}

	return
}

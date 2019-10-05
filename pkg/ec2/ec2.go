package ec2

import (
	"github.com/jdwheels/xaws/internal"
	"log"
)

func StartEC2Cluster(asgName string) bool {
	return ChangeCapacity(asgName, 1)
}

func StopEC2Cluster(asgName string) bool {
	return ChangeCapacity(asgName, 0)
}

func ChangeCapacity(asgName string, capacity int64) bool {
	_, err := internal.AutoScaleIt(asgName, capacity)
	res := err != nil
	log.Printf("Success %t", res)
	return res
}

package ec2

import (
	"github.com/jdwheels/xaws/internal"
)

func StartEC2Cluster(asgName string) bool {
	if _, err := internal.AutoScaleIt(asgName, 1); err != nil {
		return false
	} else {
		return internal.CheckIt(asgName) == nil
	}
}

func StopEC2Cluster(asgName string) bool {
	if _, err := internal.AutoScaleIt(asgName, 0); err != nil {
		return false
	} else {
		return true
	}
}

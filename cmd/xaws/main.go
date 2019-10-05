package main

import "xaws/pkg/ec2"

func main() {
	//instanceId := "i-033427252766d6858"
	//statuses, err := xaws.InstanceStatus(instanceId)
	//
	//if err != nil {
	//	return
	//}
	//
	//var status *ec2.InstanceStatus
	//if len(statuses) > 0 {
	//	status = statuses[0]
	//}
	//
	//if status != nil && *status.InstanceState.Name == ec2.InstanceStateNameRunning {
	//	log.Println("stopping...")
	//	xaws.StopInstance(instanceId, false)
	//} else {
	//	log.Println("starting...")
	//	xaws.StartInstance(instanceId, false)
	//}
	asgName := "EC2ContainerService-game-servers-2-EcsInstanceAsg-9AB2NHDSISGL"
	ec2.StopEC2Cluster(asgName)
}

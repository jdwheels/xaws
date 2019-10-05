package internal

import (
	"github.com/aws/aws-sdk-go/service/ecs"
)

func ListClusters() {
	containerService := ecs.New(GetSession())

	input := ecs.ListClustersInput{}
	req, res := containerService.ListClustersRequest(&input)

	if err := req.Send(); err != nil {
		LogAwsError(err)
		return
	}

	PrintStringArray(res.ClusterArns)
}

// Package article provides ...
package article

import (
	"context"
	"fmt"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthImpl struct{}

// TODO:1.定时检查数据库和缓存是否可用
// 2. 根据检查的结果返回服务状态
// <14-07-21, nqq> //

func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Println("this is check")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

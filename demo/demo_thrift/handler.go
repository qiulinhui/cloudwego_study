package main

import (
	"cloudwego_study/demo/demo_thrift/biz/service"
	api "cloudwego_study/demo/demo_thrift/kitex_gen/api"
	"context"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}

package app

import (
	"context"
	"log"
	pb "minx.com/demo/grpc"
)

type AppGrpc struct {
	pb.UnimplementedAppServiceServer
	AppService AppService
}

func ProviderAppGrpc (s AppService) AppGrpc {
	return AppGrpc{AppService: s}
}

func (s *AppGrpc) GetApp(ctx context.Context, in *pb.GetAppReq) (*pb.GetAppRes, error) {
	log.Printf("received: %v", in.GetId())
	return &pb.GetAppRes{App: &pb.App{ Id: 123, Download: "download", Cover: "cover"}}, nil
}

package startup

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/yongcycchen/mall-users/http_server"
	"github.com/yongcycchen/mall-users/proto/mall_users_proto/users"
	"github.com/yongcycchen/mall-users/server"
	"google.golang.org/grpc"
)

func RegisterGRPCServer(grpcServer *grpc.Server) error {
	users.RegisterUsersServiceServer(grpcServer, server.NewUsersServer())
	// users.RegisterMerchantsServiceServer(grpcServer, server.NewMerchantsServer())
	return nil
}

// RegisterGateway
func RegisterGateway(ctx context.Context, gateway *runtime.ServeMux, endPoint string, dopts []grpc.DialOption) error {
	if err := users.RegisterUsersServiceHandlerFromEndpoint(ctx, gateway, endPoint, dopts); err != nil {
		return err
	}
	if err := users.RegisterMerchantsServiceHandlerFromEndpoint(ctx, gateway, endPoint, dopts); err != nil {
		return err
	}
	return nil
}

// RegisterHttpRoute
func RegisterHttpRoute(serverMux *http.ServeMux) error {
	serverMux.HandleFunc("/swagger/", http_server.SwaggerHandler)
	return nil
}

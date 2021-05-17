package util

import (
	"context"

	"google.golang.org/grpc"
	"gitee.com/kelvins-io/kelvins/util/client_conn"
)

func GetGrpcClient(serverName string) (*grpc.ClientConn, error) {
	client, err := client_conn.NewConn(serverName)
	if err != nil {
		return nil, err
	}
	return client.GetConn(context.Background())
}

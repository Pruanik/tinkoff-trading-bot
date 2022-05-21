package grpc

import (
	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection(config *configs.Config) (grpc.ClientConnInterface, error) {
	conn, err := grpc.Dial(config.TinkoffInvestConfig.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn, nil
}

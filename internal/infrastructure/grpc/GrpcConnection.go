package grpc

import (
	"crypto/tls"
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewGrpcConnection(config *configs.Config) (grpc.ClientConnInterface, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		ServerName: config.TinkoffInvestConfig.GrpcUrl,
	})))

	token := config.TinkoffInvestConfig.SandboxToken
	if config.TinkoffInvestConfig.Mod == "production" {
		token = config.TinkoffInvestConfig.ProductionToken
	}

	opts = append(opts, grpc.WithPerRPCCredentials(AuthCredential{
		Token: token,
	}))

	addr := fmt.Sprintf("%s:%s", config.TinkoffInvestConfig.GrpcUrl, config.TinkoffInvestConfig.GrpcPort)
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

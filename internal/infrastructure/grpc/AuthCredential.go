package grpc

import "context"

type AuthCredential struct {
	Token string
}

func (ac AuthCredential) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + ac.Token,
	}, nil
}

func (ac AuthCredential) RequireTransportSecurity() bool {
	return true
}

package main

import (
	"context"
	"examples/admin/generated"
	"github.com/google/uuid"
	"github.com/iyarkov/foundation/auth"
	interceptors "github.com/iyarkov/foundation/grpc"
	"github.com/iyarkov/foundation/support"
	"github.com/iyarkov/foundation/tls"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

func main() {
	localhost := "localhost:8443"
	tlsCfg := tls.Configuration{
		CACert:     "deployment/local/ca_cert.pem",
		AppCert:    "deployment/local/admin_client_cert.pem",
		AppKey:     "deployment/local/admin_client_key.pem",
		KnownPeers: []string{"admin"},
	}
	tlsConfig, err := tlsCfg.NewCryptoTlsConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("TLS initialization failed")
	}

	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		grpc.WithChainUnaryInterceptor(
			interceptors.ClientContextId,
			interceptors.ClientAuth,
		),
	)

	conn, err := grpc.Dial(localhost, opts...)
	if err != nil {
		log.Fatal().Err(err).Msg("gRPC dial failed")
	}
	defer support.CloseWithWarning(context.Background(), conn, "")

	contextId := uuid.New()
	log.Info().Msgf("Context ID: %v", contextId)
	ctx := support.WithContextId(context.Background(), contextId.String())

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	token := auth.Token{
		AccountId: 123,
		GroupId:   21,
		Role:      auth.Admin,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	ctx = auth.WithToken(ctx, &token)

	client := generated.NewGroupsClient(conn)
	request := generated.GroupCreateRequest{
		Name: "New Group",
	}
	response, err := client.Create(ctx, &request)
	if err != nil {
		log.Fatal().Err(err).Msg("create operation failed")
	}
	if response.Code != generated.GroupModificationResponse_Ok {
		log.Fatal().Msgf("create operation failed with code %v", response.Code)
	}
	log.Info().Msgf("Group %v created, ", response.Result)
}

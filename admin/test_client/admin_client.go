package main

import (
	"context"
	"examples/admin/generated"
	"github.com/google/uuid"
	"github.com/iyarkov/foundation/auth"
	interceptors "github.com/iyarkov/foundation/grpc"
	"github.com/iyarkov/foundation/logger"
	"github.com/iyarkov/foundation/support"
	"github.com/iyarkov/foundation/telemetry"
	"github.com/iyarkov/foundation/tls"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"time"
)

func main() {
	support.AppManifest = support.Manifest{
		Instance:  "localhost-1",
		Name:      "Admin-Client",
		Version:   "1.0.1",
		Namespace: "Local",
	}
	logCfg := logger.Configuration{
		Mode:  "console",
		Level: "debug",
	}
	tlsCfg := tls.Configuration{
		CACert:     "deployment/local/ca_cert.pem",
		AppCert:    "deployment/local/admin_client_cert.pem",
		AppKey:     "deployment/local/admin_client_key.pem",
		KnownPeers: []string{"admin"},
	}
	telemetryCfg := telemetry.Configuration{
		Mode: "docker",
	}

	logger.InitLogger(&logCfg)
	ctx := support.WithContextId(context.Background(), uuid.New().String())
	log := zerolog.Ctx(ctx)
	log.Info().Msgf("configuration: %+v, %+v, %+v", logCfg, telemetryCfg, tlsCfg)

	tlsConfig, err := tlsCfg.NewCryptoTlsConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("TLS initialization failed")
	}

	telemetry.InitTelemetry(ctx, &telemetryCfg)

	ctx, span := telemetry.StartSpan(ctx, "root")
	span.SetAttributes(attribute.String("contextId", support.ContextId(ctx)))
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		grpc.WithChainUnaryInterceptor(
			interceptors.ClientContextId,
			interceptors.ClientTrace,
			interceptors.ClientAuth,
		),
	)
	localhost := "localhost:8443"
	conn, err := grpc.Dial(localhost, opts...)
	if err != nil {
		log.Fatal().Err(err).Msg("gRPC dial failed")
	}
	defer support.CloseWithWarning(context.Background(), conn, "Failed to close gRPC connection")

	log.Info().Msgf("Context ID: %v", support.ContextId(ctx))

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
		Name: "as группа крови-22",
	}
	response, err := client.Create(ctx, &request)
	span.End()
	telemetry.Flush(ctx)
	log.Info().Msgf("gRPC response %v", response)
	if err != nil {
		grpcStatus := status.Convert(err)
		if grpcStatus.Details() != nil {
			log.Error().Msgf("Error details: %v", grpcStatus.Details())
		}
		log.Fatal().Err(err).Msg("create operation failed with error")
	}
	if response.Status != nil && response.Status.Code != 0 {
		log.Fatal().Msgf("create operation failed with code %d and details: %v", response.Status.Code, response.Status.Details)
	}
	log.Info().Msgf("Group %v created, ", response.Result)
}

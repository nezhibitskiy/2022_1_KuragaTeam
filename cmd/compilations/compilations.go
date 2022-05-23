package main

import (
	"log"
	"myapp/internal/composites"
	"myapp/internal/microservices/compilations/proto"
	"net"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func main() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, err := config.Build()
	if err != nil {
		log.Fatal("zap logger build error")
	}

	logger := prLogger.Sugar()
	defer prLogger.Sync()

	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		logger.Fatal("postgres db composite failed")
	}

	lis, err := net.Listen("tcp", ":"+os.Getenv("COMPILATIONS_PORT"))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	grpcServ := grpc.NewServer()

	composite, err := composites.NewMoviesCompilationsComposite(postgresDBC, logger)
	if err != nil {
		return
	}

	proto.RegisterMovieCompilationsServer(grpcServ, composite.Service)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServ.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

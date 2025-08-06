package app

import (
	"log/slog"
	"time"

	grpcapp "sso/internal/app/grpc"
	"sso/internal/lib/kafka"
	"sso/internal/services/auth"
	"sso/internal/storage/sqlite"
)

type App struct {
	GRPCSrv       *grpcapp.App
	KafkaProducer *kafka.Producer
}

func New(
	log *slog.Logger,
	grpcProt int,
	storagePath string,
	tokenTTL time.Duration,
	producer *kafka.Producer,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL, producer)

	grpcApp := grpcapp.New(log, authService, grpcProt)

	return &App{
		GRPCSrv:       grpcApp,
		KafkaProducer: producer,
	}
}

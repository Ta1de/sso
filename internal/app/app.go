package app

import (
	"log/slog"
	"time"

	grpcapp "sso/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcProt int,
	tokenTTL time.Duration,
) *App {

	grpcApp := grpcapp.New(log, grpcProt)

	return &App{
		GRPCSrv: grpcApp,
	}
}

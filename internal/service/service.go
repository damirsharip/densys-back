package service

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/config"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
)

type Storage interface {
	ExecTX(ctx context.Context, options pgx.TxOptions, fn func(queries *pgstorage.Queries) error) error
}

type Service struct {
	cfg     config.Config
	storage Storage
}

func NewService(cfg config.Config, storage Storage) *Service {
	return &Service{
		cfg:     cfg,
		storage: storage,
	}
}

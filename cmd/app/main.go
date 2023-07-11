package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/khanfromasia/densys/admin/internal/config"
	"github.com/khanfromasia/densys/admin/internal/delivery/http"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/pkg/pgx"
	"github.com/khanfromasia/densys/admin/internal/server"
	"github.com/khanfromasia/densys/admin/internal/service"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("./config.yml"); err != nil {
		log.Fatalf("failed to read config fail %s", err.Error())
	}

	cfg := config.Get()

	pool, err := pgx.NewPool(ctx, cfg.PgxPool, cfg.Database.Dsn)

	if err != nil {
		log.Println("Failed to open a postgres pool: ", err.Error())
		return
	}

	defer pool.Close()

	storage := pgstorage.NewStorage(pool)
	svc := service.NewService(cfg, storage)

	if _, err := svc.UserGetByEmail(ctx, "admin@admin.com"); err != nil {
		user, err := svc.AdminCreate(ctx, entity.User{
			Email:    "admin@admin.com",
			Password: "admin",
		})
		log.Println("Admin created: ", user, err)
	}

	handler := http.NewHandler(svc)

	srv := server.NewHTTPServer(handler.SetupRoutes())

	go srv.Run(cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Println(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		log.Println(fmt.Sprintf("ctx.Done: %v", done))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Println("failed to shutdown server: ", err)
	}
}

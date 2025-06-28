package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"internship/internal/logg"
	"os"
	"time"
)

func ConnectDatabase() *pgxpool.Pool {
	logg.Logger.Info("Подлючение к базе данных")

	if err := godotenv.Load(); err != nil {
		logg.Logger.Error("Не удалось загрузить .env файл", zap.Error(err))
		return nil
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_SSLMODE"),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	dbpool, err := pgxpool.New(ctx, connStr)

	err = dbpool.Ping(ctx)

	if err != nil {
		logg.Logger.Error("Не удалось подключиться к базе данных: %v\n", zap.Error(err))
		return nil
	}
	logg.Logger.Info("Подключение к базе данных прошла успешно.")
	return dbpool
}

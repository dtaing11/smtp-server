package db

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func DbConnect(ctx context.Context) (*sql.DB, error) {
	instance := os.Getenv("INSTANCE_CONNECTION_NAME")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	usePrivate := os.Getenv("PRIVATE_IP") // set to "true" if you enable private IP

	if instance == "" || dbName == "" || dbUser == "" || dbPassword == "" {
		return nil, fmt.Errorf("missing DB env (INSTANCE_CONNECTION_NAME/DB_NAME/DB_USER/DB_PASSWORD)")
	}

	// Create the Cloud SQL dialer
	dialer, err := cloudsqlconn.NewDialer(ctx, cloudsqlconn.WithLazyRefresh())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}

	// Build pgx config (no host/port because connector dials by instance name)
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	pgxCfg, err := pgx.ParseConfig(connStr)
	if err != nil {
		_ = dialer.Close()
		return nil, fmt.Errorf("pgx.ParseConfig: %w", err)
	}

	// Dial options (private IP optional)
	var dialOpts []cloudsqlconn.DialOption
	if usePrivate != "" && usePrivate != "0" && usePrivate != "false" {
		dialOpts = append(dialOpts, cloudsqlconn.WithPrivateIP())
	}

	// Open database/sql DB using pgx stdlib with a custom dial function
	pgxCfg.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(ctx, instance, dialOpts...)
	}
	sqlDB := stdlib.OpenDB(*pgxCfg)

	// Pool tuning for Cloud Run (keep small to avoid exhausting DB connections)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		_ = dialer.Close()
		return nil, fmt.Errorf("db ping failed: %w", err)
	}

	return sqlDB, nil
}

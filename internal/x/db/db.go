// Package db contains shared functions for db functionality
package db

import (
	"github.com/XSAM/otelsql"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx" // crdb retries and postgres interface
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Register the Postgres driver.
	"github.com/spf13/viper"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"

	xtracing "go.hollow.sh/dnscontroller/internal/x/tracing"
)

var dbDriverName = "postgres"

// NewDB creates a db object
func NewDB(logger *zap.SugaredLogger) *sqlx.DB {
	db := sqlx.MustOpen(dbDriverName, viper.GetString("db.uri"))
	db.SetMaxOpenConns(viper.GetInt("db.connections.max_open"))
	db.SetMaxIdleConns(viper.GetInt("db.connections.max_idle"))
	db.SetConnMaxIdleTime(viper.GetDuration("db.connections.max_lifetime"))

	if err := db.Ping(); err != nil {
		logger.Fatalw("failed verifying database connection", "error", err)
	}

	return db
}

// NewDBWithTracing Creates a db object with tracing enabled
func NewDBWithTracing(logger *zap.SugaredLogger) *sqlx.DB {
	xtracing.New(viper.GetString("tracing.endpoint"), logger)

	// Register an otel sql driver
	var err error

	dbDriverName, err = otelsql.Register(dbDriverName, otelsql.WithAttributes(semconv.DBSystemCockroachdb))

	if err != nil {
		logger.Fatalw("failed initializing sql tracer", "error", err)
	}

	return NewDB(logger)
}

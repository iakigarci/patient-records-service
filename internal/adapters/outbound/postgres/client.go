package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/iakigarci/go-ddd-microservice-template/config"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/outbound/postgres/migrations"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	DB     *sqlx.DB
	logger *zap.Logger
}

func NewClient(config *config.Config, logger *zap.Logger) (*Postgres, error) {
	postgres, err := initPg(config, logger)
	if err != nil {
		logger.Error("failed to initialize postgres client", zap.Error(err))
		return nil, err
	}

	// Run migrations
	if err := migrations.RunMigrations(postgres.DB.DB); err != nil {
		logger.Error("failed to run migrations", zap.Error(err))
		return nil, err
	}

	return postgres, nil
}

func initPg(config *config.Config, logger *zap.Logger) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  config.Postgres.PoolMax,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		logger:       logger,
	}

	logger.Info("Postgres connection string", zap.String("connection_string", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DBName,
		config.Postgres.SSLMode,
	)))

	var err error
	for pg.connAttempts > 0 {
		connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Postgres.Host,
			config.Postgres.Port,
			config.Postgres.User,
			config.Postgres.Password,
			config.Postgres.DBName,
			config.Postgres.SSLMode,
		)
		pg.logger.Info("Postgres connection string", zap.String("conn_str", connectionString))
		pg.DB, err = sqlx.Open("postgres", connectionString)

		if err == nil {
			pg.DB.SetMaxOpenConns(pg.maxPoolSize)
			pg.DB.SetMaxIdleConns(pg.maxPoolSize)

			if err = pg.DB.Ping(); err == nil {
				break
			}
		}

		pg.logger.Info("Postgres is trying to connect, attempts left", zap.Int("attempts", pg.connAttempts))
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if err != nil {
		err = fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
		pg.logger.Error(err.Error())
		return nil, err
	}

	pg.logger.Info("Postgres connected successfully")

	return pg, nil
}

func (p *Postgres) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.DB.PingContext(ctx)
}

func (p *Postgres) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return p.DB.BeginTx(ctx, &sql.TxOptions{})
}

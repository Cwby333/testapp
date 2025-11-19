package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

type Config struct {
	Host string
	Port uint16
	Database string
	User string
	Password string
	MaxConns int
	MinConns int
	ConnIdle time.Duration
	ConnLife time.Duration
}

func New(ctx context.Context, c Config) (*Postgres, error) {
	const op = "./internal/infra/repository/postgresql/pg.go.New()"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d&pool_min_conns=%d", c.User, c.Password, c.Host, c.Port, c.Database, c.MaxConns, c.MinConns)
	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	fmt.Println(poolCfg.ConnString())
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {	
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Postgres{
		pool: pool,
	}, nil
}

func (p *Postgres) Insert(ctx context.Context, number int) error {
	const op = "./internal/infra/repository/postgresql/pg.go.Insert()"
	const q = "INSERT INTO testapp_table(value) VALUES ($1)"

	_, err := p.pool.Exec(ctx, q, number)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (p *Postgres) Get(ctx context.Context) ([]int, error) {
	const op = "./internal/infra/repository/postgresql/pg.go.Get()"
	const q = "SELECT value FROM testapp_table"

	rows, err := p.pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	res := make([]int, 0)
	for rows.Next() {
		var number int
		err = rows.Scan(&number)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		res = append(res, number)
	}

	return res, nil
}
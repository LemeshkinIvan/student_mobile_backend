package db

import (
	c_err "app/internal/common/error"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Password string
	UserName string
	Port     uint
	Url      string
	Pool     *pgxpool.Pool
	// logger
}

// alies
// pass == password, u == userName, p == port
func Connect(pass string, u string, p uint, url string, d_n string) (*Postgres, error) {
	conn_url := fmt.Sprintf("%s:%d", url, p)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", u, pass, conn_url, d_n)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, c_err.GetError(err.Error(), "Connect")
	}

	db := &Postgres{
		Password: pass,
		UserName: u,
		Port:     p,
		Url:      url,
		Pool:     pool,
	}

	return db, nil
}

func (db *Postgres) ExecuteQuery() {

}

func (db *Postgres) Disconnect() {
	db.Pool.Close()
}

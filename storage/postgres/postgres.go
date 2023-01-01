package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	homeDB *sqlx.DB
}

func InitDB(psqlUrl string) (psql *Postgres, err error) {

	psqlDB, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		return psql, err
	} 

	return &Postgres{
		homeDB: psqlDB,
	}, nil
}
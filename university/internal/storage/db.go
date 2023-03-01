package storage

import (
	"context"
	"github.com/Studio56School/university/internal/config"
	"github.com/jackc/pgx/v5"
	"log"
)

func ConnectDB(conf *config.Config) (*pgx.Conn, error) {

	connString := "postgres://" + conf.Username + ":" +
		conf.Password + "@" + conf.Host + ":" +
		conf.Port + "/" + conf.DBname

	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatal(err)
	}

	return conn, nil

}

package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPool() *pgxpool.Pool {

	username := "postgres"
	//this is non sensitive for now since its on local docker instances
	password := "mysecretpassword"

	url := "localhost:5432"
	dbName := "postgres"

	fullConfig := "postgres://" + username + ":" + password + "@" + url + "/" + dbName

	config, err := pgxpool.ParseConfig(fullConfig)
	if err != nil {
		log.Fatal("Can't connect to database , please check out your db settings")
	}
	config.MaxConns = 4

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Can't connect to database , please check out your db settings")
	}

	return pool

}

// func addTypes(config *pgxpool.Config) {
// 	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
// 		conn.ConnInfo().RegisterDataType(pgtype.DataType{
// 			Value: &uuid.UUID{},
// 			Name:  "uuid",
// 			OID:   pgtype.UUIDOID,
// 		})
// 		return nil
// 	}
// }

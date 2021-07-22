package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mananwalia959/go-todos-app/pkg/config"
)

func GetPool(appconfig *config.Appconfig) *pgxpool.Pool {

	username := appconfig.PostgresUsername
	//this is non sensitive for now since its on local docker instances
	password := appconfig.PostgresPassword

	url := appconfig.PostgresUrl
	dbName := appconfig.PostgresDbName

	fullConfig := "postgres://" + username + ":" + password + "@" + url + "/" + dbName

	log.Println("Trying to connect to Postgres , this might take a while ...")

	config, err := pgxpool.ParseConfig(fullConfig)
	if err != nil {
		log.Fatal("Can't connect to database , please check out your db settings")
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Can't connect to database , please check out your db settings")
	}
	log.Println("Postgres Pool Successfully Initilized")
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

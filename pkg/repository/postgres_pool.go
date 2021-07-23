package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
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
	err = migrate(pool)
	if err != nil {
		log.Fatal("Can't run migrations ")
	}

	log.Println("initial migrations successful")

	return pool

}

//this needs to be idempotent as this will execute on start each time
func migrate(pool *pgxpool.Pool) (err error) {
	tx, err := pool.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		log.Println("Can't initiate transaction")
		return err
	}

	_, err = tx.Exec(context.Background(), getQuery())

	if err != nil {
		//yuck , need a way to handle this more beautifully
		//not sure
		rollbackerr := tx.Rollback(context.Background())
		if rollbackerr != nil {
			log.Println(rollbackerr)
		}

		return err
	}

	return tx.Commit(context.Background())

}

func getQuery() string {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		email TEXT,
		UNIQUE(email)
	);

	CREATE TABLE IF NOT EXISTS todos (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT,
		description TEXT, 
		completed BOOLEAN , 
		created_on TIMESTAMP ,
		created_by UUID, 
		archived BOOLEAN DEFAULT FALSE,
		CONSTRAINT fk_user
			FOREIGN KEY(created_by) 
				REFERENCES users(id) ON DELETE CASCADE

	);

	CREATE INDEX IF NOT EXISTS todos_created_by_idx ON todos( created_by );
	`
	return query
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chelnik/sbercloudServer/config"
	"github.com/chelnik/sbercloudServer/utils"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

const (
	nameConfig    string = "/config.yaml"
	selectConfigs string = `CREATE TABLE IF NOT EXISTS configs(
								service VARCHAR(80),
								data JSON,
								version INT
							)`
	insertConfigs = `INSERT INTO configs (Service, Data, Version) VALUES (
	        $1, -- Service
			$2, -- Version
			$3 -- Data
			)`
	selectMaxConfig = `SELECT * FROM configs WHERE service=$1 ORDER BY version DESC LIMIT 1`
	deleteMaxConfig = `DELETE FROM configs WHERE service=$1 AND version=$2`
)

func main() {
	address := config.NewPointerAddress() // работа с конфигом
	err := address.LoadConfig(utils.TakeCurrentDirectory() + nameConfig)
	if err != nil {
		log.Fatalln("Error with LoadConfig", err)
	}
	// --------------открываю базу

	dsn := os.Getenv("POSTGRES_URI")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Printf("Подключение не установилось:\t %s\n", dsn)
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Подключение с базой данных установленно:\t", dsn)

	defer conn.Close(context.Background())

	if err = conn.Ping(context.Background()); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}
	// --------------------

	data := &Data{conn}
	data.InitDb()
	router := mux.NewRouter() // подключаю роутер горилла

	data.addHandlers(router)
	coveredMux := data.middlewareHandler(router)
	// fmt.Printf("server listen at http://localhost%s\n", address.Port)
	log.Printf("server listen at http://localhost%s\n", address.Port)

	err = http.ListenAndServe(address.Port, coveredMux)
	if err != nil {
		log.Fatalln("error with server")
	}
}

// InitDb создаем бд если не создана
func (d *Data) InitDb() {
	rows, err := d.conn.Query(context.Background(), selectConfigs)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

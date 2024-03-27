package cmd

import (
	"ProjectL2/develop/dev11/config"
	"ProjectL2/develop/dev11/internal/db"
	"ProjectL2/develop/dev11/internal/server"
	"fmt"

	"log"
)

// Execute запускает основные функции программы.
func Execute() {
	initializeLogger()

	serverCfg := config.GetServerConf("cfg.ini")
	dbCfg := config.GetDBConnectionConf()

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.DBName, dbCfg.Login, dbCfg.Password)

	dbConnection := db.NewConnection(dsn)

	server.Up(dbConnection, serverCfg)
}

func initializeLogger() {
	log.SetFlags(log.Ldate | log.Llongfile)
}

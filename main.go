package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/stdapps/graphql-example/delivery/graphql"
	"github.com/stdapps/graphql-example/storage"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	connstr := viper.GetString("database.url")
	postgresStorage := storage.NewPostgresStorage(storage.OpenPostgresDB(connstr))

	graphql.StartGraphqlServer(viper.GetString("server.port"), postgresStorage)

}

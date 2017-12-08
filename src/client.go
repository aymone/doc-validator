package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var client *mgo.Database

type connectionConfig struct {
	host   string
	dbName string
}

func getConnectionConfig() connectionConfig {
	host, exist := os.LookupEnv("APP_MONGO_HOST")
	if !exist {
		log.Fatal("Undefined env: APP_MONGO_HOST")
	}

	port, exist := os.LookupEnv("APP_MONGO_PORT")
	if !exist {
		log.Fatal("Undefined env: APP_MONGO_PORT")
	}

	dbName, exist := os.LookupEnv("APP_MONGO_DB_NAME")
	if !exist {
		log.Fatal("Undefined env: APP_MONGO_DB_NAME")
	}

	c := connectionConfig{
		host:   fmt.Sprintf("%s:%s", host, port),
		dbName: dbName,
	}

	return c
}

func getClient() *mgo.Database {
	if client != nil {
		return client
	}

	config := getConnectionConfig()
	session, err := mgo.Dial(config.host)
	if err != nil {
		log.Fatal(err)
	}

	client = session.DB(config.dbName)
	return client
}

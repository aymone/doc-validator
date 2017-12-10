package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

// client is a global var to store one instance from mongodb client
var client *mgo.Database

// connectionConfig is the context to run client
// host are the mongodb address
// dbName are the mongodb database
type connectionConfig struct {
	host   string
	dbName string
}

// getConnectionConfig read config from env
// This env vars are defined in host container
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

// getClient returns mongoDb client.
// If already defined, return the same instance
// if not, create a new one.
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

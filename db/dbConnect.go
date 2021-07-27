package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Connect function for connenting to DB
func Connect() (*mongo.Database, context.Context, context.CancelFunc) {
	dbUser := "nahid_Anexa"
	dbPass := "nahidForAnexa"
	dbName := "anexa_test"

	//this is mongoDB atlas connection string
	connectionString := "mongodb+srv://" + dbUser + ":" + dbPass + "@testcluster.kwwik.gcp.mongodb.net/" + dbName + "?retryWrites=true&w=majority"

	dbClient, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}

	//return db
	return dbClient.Database(dbName), ctx, cancel
}

// ###############################################################################################################
// NOT FINISHED
// ###############################################################################################################

package REST_API_MONGO

import (
	"context"

	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func ConnectDB() *mongo.Collection {
	Mongo_URI := ""
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URI))
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connected to mongoDB and pinged.")
	collection := client.Database("").Collection("")
	return collection
}

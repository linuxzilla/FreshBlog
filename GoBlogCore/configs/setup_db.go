package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(MongoUri).
			SetAuth(options.Credential{
				AuthSource: DbName,
				Username:   GetEnv("DB_USER"),
				Password:   GetEnv("DB_PWD"),
			}))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

// collection names
const (
	articles = "articles"
	authors  = "authors"
	comments = "comments"
)

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(DbName).Collection(collectionName)
	return collection
}

func InitDB(client *mongo.Client, ctx context.Context) {
	err := client.Database(DbName).CreateCollection(ctx, articles)
	if err != nil {
		log.Fatal(err)
	}
	CreateIndex(client, articles, "title", true)
	CreateIndex(client, articles, "tags", false)
	CreateIndex(client, articles, "authors", false)

	err = client.Database(DbName).CreateCollection(ctx, authors)
	if err != nil {
		log.Fatal(err)
	}
	CreateIndex(client, authors, "username", true)

	err = client.Database(DbName).CreateCollection(ctx, comments)
	if err != nil {
		log.Fatal(err)
	}
	CreateIndex(client, authors, "article_id", true)

}

func CreateIndex(client *mongo.Client, collectionName string, field string, unique bool) bool {

	// 1. Let's define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	// 2. Create the context for this operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Connect to the database and access the collection
	collection := client.Database(DbName).Collection(collectionName)

	// 4. Create a single index
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		// 5. Something went wrong, we log it and return false
		fmt.Println(err.Error())
		return false
	}

	// 6. All went well, we return true
	return true
}

package mongo_db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type MongoDB struct {
	Client             *mongo.Client
	Context            context.Context
	CategoryCollection *mongo.Collection
	CityCollection     *mongo.Collection
	VacancyCollection  *mongo.Collection
}

func NewMongoDB() *MongoDB {
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error loading .env file")
	}

	Ctx := context.TODO()
	dbHost := os.Getenv("MONGO_HOST")
	dbPort := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB")
	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	credential := options.Credential{
		Username: username,
		Password: password,
	}
	connectionURI := "mongodb://" + dbHost + ":" + dbPort + "/"
	clientOptions := options.Client().ApplyURI(connectionURI).SetAuth(credential)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database(dbName)
	CategoryCollection := db.Collection("categories")
	CityCollection := db.Collection("cities")
	VacancyCollection := db.Collection("vacancies")
	return &MongoDB{
		Client:             client,
		Context:            Ctx,
		CategoryCollection: CategoryCollection,
		CityCollection:     CityCollection,
		VacancyCollection:  VacancyCollection,
	}
}

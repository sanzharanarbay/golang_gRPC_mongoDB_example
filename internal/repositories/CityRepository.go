package repositories

import (
	"errors"
	mongo_db "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/config/mongo-db"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CityRepository struct {
	MongoDB *mongo_db.MongoDB
}

func NewCityRepository(MongoDB *mongo_db.MongoDB) *CityRepository {
	return &CityRepository{
		MongoDB: MongoDB,
	}
}

type CityRepositoryInterface interface {
	SaveCity(city *models.City) (interface{}, error)
	GetCity(ID string) (*models.City, error)
	GetAllCities() (*[]models.City, error)
	UpdateCity(city *models.City, ID string) (interface{}, error)
	DeleteCity(ID string) error
	CheckCityExistence(Label string) (*models.City, error)
}

func (c *CityRepository) SaveCity(city *models.City) (interface{}, error) {
	city.CreatedAt = time.Now()
	result, err := c.MongoDB.CityCollection.InsertOne(c.MongoDB.Context, city)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (c *CityRepository) GetCity(ID string) (*models.City, error) {
	var city models.City
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = c.MongoDB.CityCollection.FindOne(c.MongoDB.Context, bson.D{{"_id", objectId}}).Decode(&city)
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (c *CityRepository) GetAllCities() (*[]models.City, error) {
	var city models.City
	var cities []models.City

	cursor, err := c.MongoDB.CityCollection.Find(c.MongoDB.Context, bson.D{})
	if err != nil {
		defer cursor.Close(c.MongoDB.Context)
		return nil, err
	}

	for cursor.Next(c.MongoDB.Context) {
		err := cursor.Decode(&city)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return &cities, nil
}

func (c *CityRepository) UpdateCity(city *models.City, ID string) (interface{}, error) {
	city.UpdatedAt = time.Now()
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"label", city.Label}, {"total", city.Total},
		{"status", city.Status}, {"order", city.Order},
		{"updated_at", city.UpdatedAt},
	}}}

	res := c.MongoDB.CityCollection.FindOneAndUpdate(
		c.MongoDB.Context,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	var cityUpdated models.City

	if err := res.Decode(&cityUpdated); err != nil {
		return nil, err
	}

	return &cityUpdated, nil
}

func (c *CityRepository) DeleteCity(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	res, err := c.MongoDB.CityCollection.DeleteOne(c.MongoDB.Context, bson.D{{"_id", objectId}})

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (c *CityRepository) CheckCityExistence(Label string) (*models.City, error) {
	var city models.City
	filter := bson.D{{"label", Label}}
	err := c.MongoDB.CityCollection.FindOne(c.MongoDB.Context, filter).Decode(&city)
	if err != nil {
		return nil, err
	}
	return &city, nil
}

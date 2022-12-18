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

type CategoryRepository struct {
	MongoDB *mongo_db.MongoDB
}

func NewCategoryRepository(MongoDB *mongo_db.MongoDB) *CategoryRepository {
	return &CategoryRepository{
		MongoDB: MongoDB,
	}
}

type CategoryRepositoryInterface interface {
	SaveCategory(category *models.Category) (interface{}, error)
	GetCategory(ID string) (*models.Category, error)
	GetAllCategories() (*[]models.Category, error)
	UpdateCategory(category *models.Category, ID string) (interface{}, error)
	DeleteCategory(ID string) error
	CheckCategoryExistence(Label string) (*models.Category, error)
}

func (c *CategoryRepository) SaveCategory(category *models.Category) (interface{}, error) {
	category.CreatedAt = time.Now()
	result, err := c.MongoDB.CategoryCollection.InsertOne(c.MongoDB.Context, category)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (c *CategoryRepository) GetCategory(ID string) (*models.Category, error) {
	var category models.Category
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = c.MongoDB.CategoryCollection.FindOne(c.MongoDB.Context, bson.D{{"_id", objectId}}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepository) GetAllCategories() (*[]models.Category, error) {
	var category models.Category
	var categories []models.Category

	cursor, err := c.MongoDB.CategoryCollection.Find(c.MongoDB.Context, bson.D{})
	if err != nil {
		defer cursor.Close(c.MongoDB.Context)
		return nil, err
	}

	for cursor.Next(c.MongoDB.Context) {
		err := cursor.Decode(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return &categories, nil
}

func (c *CategoryRepository) UpdateCategory(category *models.Category, ID string) (interface{}, error) {
	category.UpdatedAt = time.Now()
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"label", category.Label}, {"total", category.Total},
		{"status", category.Status}, {"order", category.Order},
		{"updated_at", category.UpdatedAt},
	}}}

	res := c.MongoDB.CategoryCollection.FindOneAndUpdate(
		c.MongoDB.Context,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	var categoryUpdated models.Category

	if err := res.Decode(&categoryUpdated); err != nil {
		return nil, err
	}

	return &categoryUpdated, nil
}

func (c *CategoryRepository) DeleteCategory(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	res, err := c.MongoDB.CategoryCollection.DeleteOne(c.MongoDB.Context, bson.D{{"_id", objectId}})

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (c *CategoryRepository) CheckCategoryExistence(Label string) (*models.Category, error) {
	var category models.Category
	filter := bson.D{{"label", Label}}
	err := c.MongoDB.CategoryCollection.FindOne(c.MongoDB.Context, filter).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

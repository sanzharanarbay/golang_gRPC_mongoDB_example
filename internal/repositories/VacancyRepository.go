package repositories

import (
	"errors"
	"github.com/gosimple/slug"
	mongo_db "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/config/mongo-db"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type VacancyRepository struct {
	MongoDB *mongo_db.MongoDB
}

func NewVacancyRepository(MongoDB *mongo_db.MongoDB) *VacancyRepository {
	return &VacancyRepository{
		MongoDB: MongoDB,
	}
}

type VacancyRepositoryInterface interface {
	SaveVacancy(vacancy *models.Vacancy) (interface{}, error)
	GetVacancy(ID string) (*models.Vacancy, error)
	GetAllVacancies() (*[]models.Vacancy, error)
	UpdateVacancy(vacancy *models.Vacancy, ID string) (interface{}, error)
	DeleteVacancy(ID string) error
	CheckVacancyExistence(Title string) (*models.Vacancy, error)
}

func (v *VacancyRepository) SaveVacancy(vacancy *models.Vacancy) (interface{}, error) {
	vacancy.CreatedAt = time.Now()
	vacancy.Slug = slug.Make(vacancy.Title)
	result, err := v.MongoDB.VacancyCollection.InsertOne(v.MongoDB.Context, vacancy)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (v *VacancyRepository) GetVacancy(ID string) (*models.Vacancy, error) {
	var vacancy models.Vacancy
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	err = v.MongoDB.VacancyCollection.FindOne(v.MongoDB.Context, bson.D{{"_id", objectId}}).Decode(&vacancy)
	if err != nil {
		return nil, err
	}
	return &vacancy, nil
}

func (v *VacancyRepository) GetAllVacancies() (*[]models.Vacancy, error) {
	var vacancy models.Vacancy
	var vacancies []models.Vacancy

	cursor, err := v.MongoDB.VacancyCollection.Find(v.MongoDB.Context, bson.D{})
	if err != nil {
		defer cursor.Close(v.MongoDB.Context)
		return nil, err
	}

	for cursor.Next(v.MongoDB.Context) {
		err := cursor.Decode(&vacancy)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, vacancy)
	}

	return &vacancies, nil
}

func (v *VacancyRepository) UpdateVacancy(vacancy *models.Vacancy, ID string) (interface{}, error) {
	vacancy.UpdatedAt = time.Now()
	vacancy.Slug = slug.Make(vacancy.Title)
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"title", vacancy.Title}, {"slug", vacancy.Slug},
		{"description", vacancy.Description}, {"salary", vacancy.Salary},
		{"is_active", vacancy.IsActive}, {"is_distribute", vacancy.IsDistribute},
		{"cities", vacancy.Cities}, {"categories", vacancy.Categories},
		{"order", vacancy.Order}, {"updated_at", vacancy.UpdatedAt},
	}}}

	res := v.MongoDB.VacancyCollection.FindOneAndUpdate(
		v.MongoDB.Context,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	var vacancyUpdated models.Vacancy

	if err := res.Decode(&vacancyUpdated); err != nil {
		return nil, err
	}

	return &vacancyUpdated, nil
}

func (v *VacancyRepository) DeleteVacancy(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	res, err := v.MongoDB.VacancyCollection.DeleteOne(v.MongoDB.Context, bson.D{{"_id", objectId}})

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (v *VacancyRepository) CheckVacancyExistence(Title string) (*models.Vacancy, error) {
	var vacancy models.Vacancy
	filter := bson.D{{"title", Title}}
	err := v.MongoDB.VacancyCollection.FindOne(v.MongoDB.Context, filter).Decode(&vacancy)
	if err != nil {
		return nil, err
	}
	return &vacancy, nil
}

package services

import (
	"errors"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/repositories"
)

type CityService struct {
	cityRepository *repositories.CityRepository
}

func NewCityService(cityRepository *repositories.CityRepository) *CityService {
	return &CityService{
		cityRepository: cityRepository,
	}
}

func (c *CityService) Insert(city *models.City) (interface{}, error) {
	check, _ := c.cityRepository.CheckCityExistence(city.Label)
	if check != nil {
		return nil, errors.New("The City already exists")
	}
	state, err := c.cityRepository.SaveCity(city)
	return state, err

}

func (c *CityService) Get(ID string) (*models.City, error) {
	state, err := c.cityRepository.GetCity(ID)
	return state, err
}

func (c *CityService) List() (*[]models.City, error) {
	state, err := c.cityRepository.GetAllCities()
	return state, err
}

func (c *CityService) Update(city *models.City, ID string) (interface{}, error) {
	state, err := c.cityRepository.UpdateCity(city, ID)
	return state, err
}

func (c *CityService) Delete(ID string) error {
	err := c.cityRepository.DeleteCity(ID)
	return err
}

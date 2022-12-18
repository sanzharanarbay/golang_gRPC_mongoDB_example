package services

import (
	"errors"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/repositories"
)

type VacancyService struct {
	vacancyRepository  *repositories.VacancyRepository
	categoryRepository *repositories.CategoryRepository
	cityRepository     *repositories.CityRepository
}

func NewVacancyService(vacancyRepository *repositories.VacancyRepository, categoryRepository *repositories.CategoryRepository,
	cityRepository *repositories.CityRepository) *VacancyService {
	return &VacancyService{
		vacancyRepository:  vacancyRepository,
		categoryRepository: categoryRepository,
		cityRepository:     cityRepository,
	}
}

func (v *VacancyService) Insert(requestBody *models.VacancyRequestBody) (interface{}, error) {
	var categories []*models.Category
	var cities []*models.City
	var vacancy models.Vacancy

	for _, element := range requestBody.Categories {
		categoryObj, _ := v.categoryRepository.GetCategory(element)
		if categoryObj == nil {
			return nil, errors.New("no category with that Id exists")
		}
		categories = append(categories, categoryObj)
	}

	for _, element := range requestBody.Cities {
		cityObj, _ := v.cityRepository.GetCity(element)
		if cityObj == nil {
			return nil, errors.New("no city with that Id exists")
		}
		cities = append(cities, cityObj)
	}

	check, _ := v.vacancyRepository.CheckVacancyExistence(requestBody.Title)
	if check != nil {
		return nil, errors.New("the Vacancy with this title already exists")
	}

	vacancy.Title = requestBody.Title
	vacancy.Description = requestBody.Description
	vacancy.Salary = requestBody.Salary
	vacancy.IsActive = requestBody.IsActive
	vacancy.IsDistribute = requestBody.IsDistribute
	vacancy.Order = requestBody.Order
	vacancy.Categories = categories
	vacancy.Cities = cities

	state, err := v.vacancyRepository.SaveVacancy(&vacancy)
	return state, err

}

func (v *VacancyService) Get(ID string) (*models.Vacancy, error) {
	state, err := v.vacancyRepository.GetVacancy(ID)
	return state, err
}

func (v *VacancyService) List() (*[]models.Vacancy, error) {
	state, err := v.vacancyRepository.GetAllVacancies()
	return state, err
}

func (v *VacancyService) Update(requestBody *models.VacancyRequestBody, ID string) (interface{}, error) {
	var categories []*models.Category
	var cities []*models.City
	var vacancy models.Vacancy

	for _, element := range requestBody.Categories {
		categoryObj, _ := v.categoryRepository.GetCategory(element)
		if categoryObj == nil {
			return nil, errors.New("no category with that Id exists")
		}
		categories = append(categories, categoryObj)
	}

	for _, element := range requestBody.Cities {
		cityObj, _ := v.cityRepository.GetCity(element)
		if cityObj == nil {
			return nil, errors.New("no city with that Id exists")
		}
		cities = append(cities, cityObj)
	}

	vacancy.Title = requestBody.Title
	vacancy.Description = requestBody.Description
	vacancy.Salary = requestBody.Salary
	vacancy.IsActive = requestBody.IsActive
	vacancy.IsDistribute = requestBody.IsDistribute
	vacancy.Order = requestBody.Order
	vacancy.Categories = categories
	vacancy.Cities = cities

	state, err := v.vacancyRepository.UpdateVacancy(&vacancy, ID)
	return state, err

}

func (v *VacancyService) Delete(ID string) error {
	err := v.vacancyRepository.DeleteVacancy(ID)
	return err
}

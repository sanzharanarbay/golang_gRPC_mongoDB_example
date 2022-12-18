package services

import (
	"errors"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/repositories"
)

type CategoryService struct {
	categoryRepository *repositories.CategoryRepository
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryService) Insert(category *models.Category) (interface{}, error) {
	check, _ := c.categoryRepository.CheckCategoryExistence(category.Label)
	if check != nil {
		return nil, errors.New("The Category already exists")
	}
	state, err := c.categoryRepository.SaveCategory(category)
	return state, err

}

func (c *CategoryService) Get(ID string) (*models.Category, error) {
	state, err := c.categoryRepository.GetCategory(ID)
	return state, err
}

func (c *CategoryService) List() (*[]models.Category, error) {
	state, err := c.categoryRepository.GetAllCategories()
	return state, err
}

func (c *CategoryService) Update(category *models.Category, ID string) (interface{}, error) {
	state, err := c.categoryRepository.UpdateCategory(category, ID)
	return state, err
}

func (c *CategoryService) Delete(ID string) error {
	err := c.categoryRepository.DeleteCategory(ID)
	return err
}

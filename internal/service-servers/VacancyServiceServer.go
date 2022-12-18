package service_servers

import (
	"context"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	vacancyServer "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/pb"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/repositories"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/requests"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type VacancyServiceServer struct {
}

var vacancyRepository = repositories.NewVacancyRepository(mongoDB)
var vacancyService = services.NewVacancyService(vacancyRepository, categoryRepository, cityRepository)

func getVacancyData(data *models.Vacancy) *vacancyServer.Vacancy {
	return &vacancyServer.Vacancy{
		Id:           data.Id.Hex(),
		Title:        data.Title,
		Slug:         data.Slug,
		Description:  data.Description,
		Salary:       data.Salary,
		IsActive:     data.IsActive,
		IsDistribute: data.IsDistribute,
		Order:        data.Order,
		CreatedAt:    data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func getVacancyCity(data *models.City) *vacancyServer.City {
	return &vacancyServer.City{
		Id:        data.Id.Hex(),
		Label:     data.Label,
		Total:     data.Total,
		Status:    data.Status,
		Order:     data.Order,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func getVacancyCategory(data *models.Category) *vacancyServer.Category {
	return &vacancyServer.Category{
		Id:        data.Id.Hex(),
		Label:     data.Label,
		Total:     data.Total,
		Status:    data.Status,
		Order:     data.Order,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (v *VacancyServiceServer) ReadVacancy(ctx context.Context, req *vacancyServer.ReadVacancyRequest) (*vacancyServer.ReadVacancyResponse, error) {
	vacancy, _ := vacancyService.Get(req.GetId())

	var cities []*vacancyServer.City
	var categories []*vacancyServer.Category

	if vacancy == nil {
		err := status.Error(codes.NotFound, "The vacancy was not found")
		return nil, err
	}

	for _, element := range vacancy.Cities {
		response := getVacancyCity(element)
		cities = append(cities, response)
	}

	for _, element := range vacancy.Categories {
		response := getVacancyCategory(element)
		categories = append(categories, response)
	}

	vacancyData := getVacancyData(vacancy)
	vacancyData.Cities = cities
	vacancyData.Categories = categories

	response := &vacancyServer.ReadVacancyResponse{
		Vacancy: vacancyData,
	}
	return response, nil
}

func (v *VacancyServiceServer) ListVacancies(ctx context.Context, req *vacancyServer.ListVacancyRequest) (*vacancyServer.ListVacancyResponse, error) {
	vacancyList, _ := vacancyService.List()

	var vacancies []*vacancyServer.Vacancy

	for _, element := range *vacancyList {

		var cities []*vacancyServer.City
		var categories []*vacancyServer.Category

		for _, item := range element.Cities {
			response := getVacancyCity(item)
			cities = append(cities, response)
		}
		for _, item := range element.Categories {
			response := getVacancyCategory(item)
			categories = append(categories, response)
		}
		response := getVacancyData(&element)
		response.Cities = cities
		response.Categories = categories
		vacancies = append(vacancies, response)
	}

	response := &vacancyServer.ListVacancyResponse{
		Vacancies: vacancies,
	}

	return response, nil
}

func (v *VacancyServiceServer) CreateVacancy(ctx context.Context, req *vacancyServer.CreateVacancyReq) (*vacancyServer.CreateVacancyResponse, error) {
	vacancyReq := models.VacancyRequestBody{
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		Salary:       req.GetSalary(),
		IsActive:     req.GetIsActive(),
		IsDistribute: req.GetIsDistribute(),
		Order:        req.GetOrder(),
		Categories:   req.GetCategories(),
		Cities:       req.GetCities(),
	}

	validErrs := requests.NewVacancyRequest(&vacancyReq).ValidateVacancy()

	if len(validErrs) > 0 {
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	_, err := vacancyService.Insert(&vacancyReq)

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("Vacancy Saved Successfully")
	}

	response := &vacancyServer.CreateVacancyResponse{
		Status:  1,
		Message: "Vacancy Saved Successfully",
	}
	return response, nil
}

func (v *VacancyServiceServer) UpdateVacancy(ctx context.Context, req *vacancyServer.UpdateVacancyReq) (*vacancyServer.UpdateVacancyResponse, error) {
	vacancyReq := models.VacancyRequestBody{
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		Salary:       req.GetSalary(),
		IsActive:     req.GetIsActive(),
		IsDistribute: req.GetIsDistribute(),
		Order:        req.GetOrder(),
		Categories:   req.GetCategories(),
		Cities:       req.GetCities(),
	}

	vacancy, _ := vacancyService.Get(req.GetId())

	if vacancy == nil {
		err := status.Error(codes.NotFound, "The vacancy was not found")
		return nil, err
	}

	validErrs := requests.NewVacancyRequest(&vacancyReq).ValidateVacancy()

	if len(validErrs) > 0 {
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	_, err := vacancyService.Update(&vacancyReq, req.GetId())

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("Vacancy Updated Successfully")
	}

	response := &vacancyServer.UpdateVacancyResponse{
		Status:  1,
		Message: "Vacancy Updated Successfully",
	}
	return response, nil
}

func (v *VacancyServiceServer) DeleteVacancy(ctx context.Context, req *vacancyServer.DeleteVacancyRequest) (*vacancyServer.DeleteVacancyResponse, error) {

	vacancy, _ := vacancyService.Get(req.GetId())

	if vacancy == nil {
		err := status.Error(codes.NotFound, "The vacancy was not found")
		return nil, err
	}

	err := vacancyService.Delete(req.GetId())

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("Vacancy Deleted Successfully")
	}

	response := &vacancyServer.DeleteVacancyResponse{
		Status:  1,
		Message: "Vacancy Deleted Successfully",
	}
	return response, nil
}

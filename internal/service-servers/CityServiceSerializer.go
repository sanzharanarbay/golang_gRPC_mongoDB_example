package service_servers

import (
	"context"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	cityServer "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/pb"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/repositories"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/requests"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type CityServiceServer struct {
}

var cityRepository = repositories.NewCityRepository(mongoDB)
var cityService = services.NewCityService(cityRepository)

func getCityData(data *models.City) *cityServer.City {
	return &cityServer.City{
		Id:        data.Id.Hex(),
		Label:     data.Label,
		Total:     data.Total,
		Status:    data.Status,
		Order:     data.Order,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (c *CityServiceServer) ReadCity(ctx context.Context, req *cityServer.ReadCityRequest) (*cityServer.ReadCityResponse, error) {
	city, _ := cityService.Get(req.GetId())

	if city == nil {
		err := status.Error(codes.NotFound, "The city was not found")
		return nil, err
	}

	response := &cityServer.ReadCityResponse{
		City: getCityData(city),
	}
	return response, nil
}

func (c *CityServiceServer) ListCities(ctx context.Context, req *cityServer.ListCityRequest) (*cityServer.ListCityResponse, error) {
	cityList, _ := cityService.List()

	var cities []*cityServer.City

	for _, element := range *cityList {
		response := getCityData(&element)
		cities = append(cities, response)
	}

	response := &cityServer.ListCityResponse{
		Cities: cities,
	}

	return response, nil
}

func (c *CityServiceServer) CreateCity(ctx context.Context, req *cityServer.CreateCityReq) (*cityServer.CreateCityResponse, error) {
	cityReq := models.City{
		Label: req.GetLabel(),
		Order: req.GetOrder(),
	}

	validErrs := requests.NewCityRequest(&cityReq).ValidateCity()

	if len(validErrs) > 0 {
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	cityReq.Status = true

	_, err := cityService.Insert(&cityReq)

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("City Saved Successfully")
	}

	response := &cityServer.CreateCityResponse{
		Status:  1,
		Message: "City Saved Successfully",
	}
	return response, nil
}

func (c *CityServiceServer) UpdateCity(ctx context.Context, req *cityServer.UpdateCityReq) (*cityServer.UpdateCityResponse, error) {
	cityReq := models.City{
		Label: req.GetLabel(),
		Order: req.GetOrder(),
	}

	city, _ := cityService.Get(req.GetId())

	if city == nil {
		err := status.Error(codes.NotFound, "The city was not found")
		return nil, err
	}

	validErrs := requests.NewCityRequest(&cityReq).ValidateCity()

	if len(validErrs) > 0 {
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	cityReq.Status = true

	_, err := cityService.Update(&cityReq, req.GetId())

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("City Updated Successfully")
	}

	response := &cityServer.UpdateCityResponse{
		Status:  1,
		Message: "City Updated Successfully",
	}
	return response, nil
}

func (c *CityServiceServer) DeleteCity(ctx context.Context, req *cityServer.DeleteCityRequest) (*cityServer.DeleteCityResponse, error) {

	city, _ := cityService.Get(req.GetId())

	if city == nil {
		err := status.Error(codes.NotFound, "The city was not found")
		return nil, err
	}

	err := cityService.Delete(req.GetId())

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("City Deleted Successfully")
	}

	response := &cityServer.DeleteCityResponse{
		Status:  1,
		Message: "City Deleted Successfully",
	}
	return response, nil
}

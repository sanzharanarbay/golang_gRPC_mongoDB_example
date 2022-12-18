package service_servers

import (
	"context"
	mongo_db "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/config/mongo-db"
	categoryServer "github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/pb"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/repositories"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/requests"
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type CategoryServiceServer struct {
}

var mongoDB = mongo_db.NewMongoDB()
var categoryRepository = repositories.NewCategoryRepository(mongoDB)
var categoryService = services.NewCategoryService(categoryRepository)

func getCategoryData(data *models.Category) *categoryServer.Category {
	return &categoryServer.Category{
		Id:        data.Id.Hex(),
		Label:     data.Label,
		Total:     data.Total,
		Status:    data.Status,
		Order:     data.Order,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (c *CategoryServiceServer) ReadCategory(ctx context.Context, req *categoryServer.ReadCategoryRequest) (*categoryServer.ReadCategoryResponse, error) {
	category, _ := categoryService.Get(req.GetId())

	if category == nil {
		err := status.Error(codes.NotFound, "The category was not found")
		return nil, err
	}

	response := &categoryServer.ReadCategoryResponse{
		Category: getCategoryData(category),
	}
	return response, nil
}

func (c *CategoryServiceServer) ListCategories(ctx context.Context, req *categoryServer.ListCategoryRequest) (*categoryServer.ListCategoryResponse, error) {
	categoryList, _ := categoryService.List()

	var categories []*categoryServer.Category

	for _, element := range *categoryList {
		response := getCategoryData(&element)
		categories = append(categories, response)
	}

	response := &categoryServer.ListCategoryResponse{
		Categories: categories,
	}

	return response, nil
}

func (c *CategoryServiceServer) CreateCategory(ctx context.Context, req *categoryServer.CreateCategoryReq) (*categoryServer.CreateCategoryResponse, error) {
	categoryReq := models.Category{
		Label: req.GetLabel(),
		Order: req.GetOrder(),
	}

	validErrs := requests.NewCategoryRequest(&categoryReq).ValidateCategory()

	if len(validErrs) > 0 {
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	categoryReq.Status = true

	_, err := categoryService.Insert(&categoryReq)

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("Category Saved Successfully")
	}

	response := &categoryServer.CreateCategoryResponse{
		Status:  1,
		Message: "Category Saved Successfully",
	}
	return response, nil
}

func (c *CategoryServiceServer) UpdateCategory(ctx context.Context, req *categoryServer.UpdateCategoryReq) (*categoryServer.UpdateCategoryResponse, error) {
	categoryReq := models.Category{
		Label: req.GetLabel(),
		Order: req.GetOrder(),
	}

	category, _ := categoryService.Get(req.GetId())

	if category == nil {
		err := status.Error(codes.NotFound, "The category was not found")
		return nil, err
	}

	validErrs := requests.NewCategoryRequest(&categoryReq).ValidateCategory()

	if len(validErrs) > 0 {
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	categoryReq.Status = true

	_, err := categoryService.Update(&categoryReq, req.GetId())

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("Category Updated Successfully")
	}

	response := &categoryServer.UpdateCategoryResponse{
		Status:  1,
		Message: "Category Updated Successfully",
	}
	return response, nil
}

func (c *CategoryServiceServer) DeleteCategory(ctx context.Context, req *categoryServer.DeleteCategoryRequest) (*categoryServer.DeleteCategoryResponse, error) {

	category, _ := categoryService.Get(req.GetId())

	if category == nil {
		err := status.Error(codes.NotFound, "The category was not found")
		return nil, err
	}

	err := categoryService.Delete(req.GetId())

	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	} else {
		log.Println("Category Deleted Successfully")
	}

	response := &categoryServer.DeleteCategoryResponse{
		Status:  1,
		Message: "Category Deleted Successfully",
	}
	return response, nil
}

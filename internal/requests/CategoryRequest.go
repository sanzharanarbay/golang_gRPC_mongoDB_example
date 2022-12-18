package requests

import (
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"net/url"
)

type CategoryRequest struct {
	model *models.Category
}

func NewCategoryRequest(model *models.Category) *CategoryRequest {
	return &CategoryRequest{
		model: model,
	}
}

func (c *CategoryRequest) ValidateCategory() url.Values {
	errs := url.Values{}

	if c.model.Label == "" {
		errs.Add("label", "The label field is required!")
	}

	if len(c.model.Label) < 3 || len(c.model.Label) > 255 {
		errs.Add("label", "The label field must be between 3-255 chars!")
	}

	if c.model.Order == 0 {
		errs.Add("order", "The order field is required!")
	}

	return errs
}
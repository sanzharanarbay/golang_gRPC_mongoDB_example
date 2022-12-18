package requests

import (
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"net/url"
)

type CityRequest struct {
	model *models.City
}

func NewCityRequest(model *models.City) *CityRequest {
	return &CityRequest{
		model: model,
	}
}

func (c *CityRequest) ValidateCity() url.Values {
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

package requests

import (
	"github.com/sanzharanarbay/golang_gRPC_mongoDB_example/internal/models"
	"net/url"
)

type VacancyRequest struct {
	model *models.VacancyRequestBody
}

func NewVacancyRequest(model *models.VacancyRequestBody) *VacancyRequest {
	return &VacancyRequest{
		model: model,
	}
}

func (v *VacancyRequest) ValidateVacancy() url.Values {
	errs := url.Values{}

	if v.model.Title == "" {
		errs.Add("title", "The title field is required!")
	}

	if len(v.model.Title) < 3 || len(v.model.Title) > 255 {
		errs.Add("label", "The label field must be between 3-255 chars!")
	}

	if v.model.Description == "" {
		errs.Add("description", "The description field is required!")
	}

	if len(v.model.Description) < 5 || len(v.model.Title) > 500 {
		errs.Add("description", "The description field must be between 3-255 chars!")
	}

	if len(v.model.Cities) == 0 {
		errs.Add("cities", "The cities field is required!")
	}

	if len(v.model.Categories) == 0 {
		errs.Add("categories", "The categories field is required")
	}

	return errs
}

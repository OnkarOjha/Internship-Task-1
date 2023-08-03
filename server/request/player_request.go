package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type PlayerGetRequest struct {
	ID int `json:"id"`
}

type PlayerUpdateRequest struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (p PlayerUpdateRequest) Validate() error {
	return validation.ValidateStruct(&p,
		// Name is required and should not exceed 15 characters
		validation.Field(&p.Name, validation.Required, validation.Length(1, 15)),
		// Score is required and should be greater than or equal to 0
		validation.Field(&p.Score, validation.Required, validation.Min(0)),
	)
}

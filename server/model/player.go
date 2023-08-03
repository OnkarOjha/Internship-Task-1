package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Player represents the player attributes in the database
type Player struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Score   int    `json:"score"`
}

func (p Player) Validate() error {
	return validation.ValidateStruct(&p,
		// ID should be greater than 0 (assuming ID is a primary key)
		validation.Field(&p.ID, validation.Required, validation.Min(1)),
		// Name is required and should not exceed 15 characters
		validation.Field(&p.Name, validation.Required, validation.Length(1, 15)),
		// Country is required and should be a two-letter uppercase code
		validation.Field(&p.Country, validation.Required, validation.Length(2, 2), is.UpperCase),
		// Score is required and should be greater than or equal to 0
		validation.Field(&p.Score, validation.Required, validation.Min(0)),
	)
}

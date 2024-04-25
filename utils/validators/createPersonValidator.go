package validators

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

// PersonRequest represents the structure of the request body for creating a person
type PersonRequest struct {
    Name    string   `json:"name" validate:"required"`
    Age     int      `json:"age" validate:"required"`
    Hobbies []string `json:"hobbies" validate:"required"`
}

// ValidatePersonRequest validates the request body for creating a person
func ValidatePersonRequest(c *gin.Context) (*PersonRequest, error) {
    var req PersonRequest

    // Bind JSON data to PersonRequest struct
    if err := c.BindJSON(&req); err != nil {
        return nil, err
    }

    // Create a new validator instance
    validate := validator.New()

    // Register custom validation rules if any

    // Validate the request body
    if err := validate.Struct(req); err != nil {
        return nil, err
    }

    return &req, nil
}

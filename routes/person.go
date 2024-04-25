package routes

import (
	"GoCrudChallange/initializers/db"
	"GoCrudChallange/utils/validators"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PersonRoutes() {
	r := gin.Default()
    r.Use(cors.Default())
	// Handle requests to non-existing endpoints
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint not found"})
	})

	// Handle internal server errors
	r.Use(func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{"error": "Internal server error"})
			}
		}()
		c.Next()
	})

	// Get all persons
	r.GET("/person", func(c *gin.Context) {
		// Retrieve persons from in-memory database
		persons := db.Persons

		// Return the list of persons as JSON response
		c.JSON(200, persons)
	})

	// Get a specific person
	r.GET("/person/:personId", func(c *gin.Context) {
		// Retrieve persons from in-memory database
		persons := db.Persons

		// Extract person ID from URL parameter
		personID := c.Param("personId")

		// Find the person with the given ID
		var targetPerson map[string]interface{}
		for _, person := range persons {
			if person["id"] == personID {
				targetPerson = person
				break
			}
		}

		// Check if the person is found
		if targetPerson != nil {
			// Return the person as JSON response
			c.JSON(200, gin.H{
				"isSuccess": true,
				"person":    targetPerson,
			})
		} else {
			// If person not found, return appropriate message
			c.JSON(404, gin.H{
				"isSuccess": false,
				"message":   "Person not found",
			})
		}
	})

	// Create a new person

	r.POST("/person", func(c *gin.Context) {
		// Validate the request body
		req, err := validators.ValidatePersonRequest(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Generate UUID for the person's ID
		personID := uuid.New().String()

		// Create a map to represent the new person
		newPerson := map[string]interface{}{
			"id":      personID,
			"name":    req.Name,
			"age":     req.Age,
			"hobbies": req.Hobbies,
		}

		// Add the new person to the database
		db.Persons = append(db.Persons, newPerson)

		// Return success message
		c.JSON(200, gin.H{
			"isSuccess": true,
			"message":   "Person created successfully",
		})
	})

	// Update an existing person
	r.PUT("/person/:personId", func(c *gin.Context) {
		// Extract person ID from URL parameter
		personID := c.Param("personId")

		// Retrieve persons from in-memory database
		persons := db.Persons

		// Find the index of the person with the given ID
		var index = -1
		for i, person := range persons {
			if person["id"] == personID {
				index = i
				break
			}
		}

		// Check if the person is found
		if index != -1 {
			// Update the person data
			var updatedPerson map[string]interface{}
			if err := c.BindJSON(&updatedPerson); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			// Update the person in the database
            updatedPerson["id"] = personID
			db.Persons[index] = updatedPerson

			// Return success message
			c.JSON(200, gin.H{
				"isSuccess": true,
				"message":   "Person updated successfully",
			})
		} else {
			// If person not found, return appropriate message
			c.JSON(404, gin.H{
				"isSuccess": false,
				"message":   "Person not found",
			})
		}
	})

	// Delete an existing person
	r.DELETE("/person/:personId", func(c *gin.Context) {
		// Extract person ID from URL parameter
		personID := c.Param("personId")

		// Retrieve persons from in-memory database
		persons := db.Persons

		// Find the index of the person with the given ID
		var index = -1
		for i, person := range persons {
			if person["id"] == personID {
				index = i
				break
			}
		}

		// Check if the person is found
		if index != -1 {
			// Remove the person from the database
			db.Persons = append(persons[:index], persons[index+1:]...)

			// Return success message
			c.JSON(200, gin.H{
				"isSuccess": true,
				"message":   "Person deleted successfully",
			})
		} else {
			// If person not found, return appropriate message
			c.JSON(404, gin.H{
				"isSuccess": false,
				"message":   "Person not found",
			})
		}
	})

	r.Run() // Run the Gin router
}

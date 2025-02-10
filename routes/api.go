package routes

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rossmoney/classified-listings-go/api/responses"
)

// Listing represents data about a classified listing.
type Listing struct {
	// The ID of the listing
	//
	// example: 1
	ID int `json:"id"`
	// The title of the listing
	// required: true
	// example: Red House
	Title string `json:"title"`
	// The description of the listing
	// required: true
	// example: This is my first listing
	Description string `json:"description"`
	// The price of the listing
	// required: true
	// example: 100.00
	Price float64 `json:"price,string,omitempty"`
	// The category of the listing
	// required: true
	// example: Property
	Category string `json:"category"`
	// The date the listing was posted
	// required: true
	// example: 2025-01-01T00:00:00Z
	DatePosted string `json:"date_posted"`
	// Whether the listing is active
	// example: true
	Active bool `json:"active"`
}

// listing slice to seed listing data.
var listings = []Listing{
	{ID: 1, Title: "Red House", Description: "a Red House", Price: 560000.99, Category: "Property", DatePosted: time.Now().Format(time.RFC3339), Active: true},
	{ID: 2, Title: "Vauxhall Corza", Description: "Vauxhall Corza for sale", Price: 175.99, Category: "Vehicle", DatePosted: time.Now().Format(time.RFC3339), Active: false},
	{ID: 3, Title: "A Blue House", Description: "blue house for sale", Price: 390000.99, Category: "Property", DatePosted: time.Now().Format(time.RFC3339), Active: true},
}

// @BasePath /listings
// PingExample godoc
// @Summary get listings
// @Schemes
// @Description GetListings responds with the list of all listings as JSON.
// @Tag listings
// @Accept json
// @Produce json
// @Success      200  {object}  responses.SuccessResponse
// @Router /listings [get]
func (l *Listing) GetListings(c *gin.Context) {
	title := c.Query("title")
	category := c.Query("category")

	if title == "" && category == "" {
		c.IndentedJSON(http.StatusOK, &responses.SuccessResponse{
			Status: http.StatusOK,
			Data:   listings,
		})
		return
	}

	out := []Listing{}

	if title != "" {
		for _, l := range listings {
			if strings.Contains(l.Title, title) {
				out = append(out, l)
			}
		}
	}

	if category != "" {
		for _, l := range listings {
			if l.Category == category {
				out = append(out, l)
			}
		}
	}

	message := ""
	if len(out) == 0 {
		message = "No listings found!"
	}

	c.IndentedJSON(http.StatusOK, &responses.SuccessResponse{
		Status:  http.StatusOK,
		Data:    out,
		Message: message,
	})
}

// @BasePath /listings
// PingExample godoc
// @Summary create listing
// @Schemes
// @Description CreateListing adds a listing from JSON received in the request body.
// @Tag listings
// @Accept json
// @Produce json
// @Param listing body routes.Listing true "Listing data to create"
// @Success      201  {object}  responses.SuccessResponse
// @Failure      400  {object}  responses.ErrorMessageResponse
// @Failure      400  {object}  responses.ValidationErrorResponse
// @Router /listings [post]
func (l *Listing) CreateListing(c *gin.Context) {
	var newListing Listing

	// Call BindJSON to bind the received JSON to
	// newListing.
	if err := c.BindJSON(&newListing); err != nil {
		if strings.Contains(err.Error(), "into float64") {
			validationErrorResponse := &responses.ValidationErrorResponse{}
			validationErrorResponse.Errors = append(validationErrorResponse.Errors, responses.Error{FieldName: "price", Message: "Price must be a decimal number."})
			validationErrorResponse.Status = http.StatusBadRequest
			c.IndentedJSON(http.StatusBadRequest, validationErrorResponse)
			return
		}
		c.IndentedJSON(http.StatusBadRequest, &responses.ErrorMessageResponse{
			Message: "JSON decoding failed",
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	errors := &responses.ValidationErrorResponse{}
	errors.Status = http.StatusBadRequest

	if newListing.Title == "" {
		errors.Errors = append(errors.Errors, responses.Error{FieldName: "title", Message: "Title is required"})
	}

	if newListing.Description == "" {
		errors.Errors = append(errors.Errors, responses.Error{FieldName: "description", Message: "Description is required"})
	}

	if newListing.Category == "" {
		errors.Errors = append(errors.Errors, responses.Error{FieldName: "category", Message: "Category is required"})
	}

	if newListing.Price == 0 {
		errors.Errors = append(errors.Errors, responses.Error{FieldName: "price", Message: "Price is required"})
	}

	datePosted, err := time.Parse(time.RFC3339, newListing.DatePosted)
	if err == nil {
		newListing.DatePosted = datePosted.Format(time.RFC3339)
	}
	if datePosted.Year() == 1 {
		errors.Errors = append(errors.Errors, responses.Error{FieldName: "date_posted", Message: "Date Posted is required"})
	}

	if len(errors.Errors) > 0 {
		c.IndentedJSON(http.StatusBadRequest, errors)
		return
	}

	newListing.ID = len(listings) + 1

	// Add the new listing to the slice.
	listings = append(listings, newListing)
	c.IndentedJSON(http.StatusCreated, &responses.SuccessResponse{
		Message: "New listing saved!",
		Status:  http.StatusCreated,
		Data:    newListing,
	})
}

// @BasePath /listings/{id}
// PingExample godoc
// @Summary get listing
// @Schemes
// @Description GetListing locates the listing whose ID value matches the id parameter sent by the client, then returns that listing as a response.
// @Tag listings
// @Param id path int true "Listing ID"
// @Param title query string false "search by title"
// @Param category query string false "filter by category"
// @Accept json
// @Produce json
// @Success      200  {object}  responses.SuccessResponse
// @Failure      404  {object}  responses.ErrorMessageResponse "listing not found"
// @Router /listings/{id} [get]
func (l *Listing) GetListing(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		// Loop over the slice of listings, looking for
		// a listing whose ID value matches the parameter.
		for _, listing := range listings {
			if listing.ID == id {
				c.IndentedJSON(http.StatusOK, &responses.SuccessResponse{Data: listing, Status: http.StatusOK})
				return
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, &responses.ErrorMessageResponse{
		Message: "listing not found",
		Status:  http.StatusNotFound,
		Error:   err.Error(),
	})
}

// @BasePath /listings/{id}
// PingExample godoc
// @Summary update listing
// @Schemes
// @Description UpdateListing updates the listing whose ID value matches the id parameter sent by the client.
// @Tags listings
// @Param id path int true "Listing ID"
// @Param listing body routes.Listing true "Listing data to update"
// @Accept json
// @Produce json
// @Success      200  {object}  responses.SuccessResponse
// @Failure      400  {object}  responses.ErrorMessageResponse
// @Failure      400  {object}  responses.ValidationErrorResponse
// @Failure      404  {object}  responses.ErrorMessageResponse
// @Router /listings/{id} [put]
func (l *Listing) UpdateListing(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		// Loop over the slice of listings, looking for
		// a listing whose ID value matches the parameter.
		for i, listing := range listings {
			if listing.ID == id {
				updatedListing := listings[i]

				// Call BindJSON to bind the received JSON to
				// updatedListing.
				if err := c.BindJSON(&updatedListing); err != nil {
					c.IndentedJSON(http.StatusBadRequest, &responses.ErrorMessageResponse{
						Message: "JSON decoding failed",
						Error:   err.Error(),
						Status:  http.StatusBadRequest,
					})
					return
				}

				errors := &responses.ValidationErrorResponse{}
				errors.Status = http.StatusBadRequest

				if updatedListing.Title == "" {
					errors.Errors = append(errors.Errors, responses.Error{FieldName: "title", Message: "Title is required"})
				}

				if updatedListing.Description == "" {
					errors.Errors = append(errors.Errors, responses.Error{FieldName: "description", Message: "Description is required"})
				}

				if updatedListing.Category == "" {
					errors.Errors = append(errors.Errors, responses.Error{FieldName: "category", Message: "Category is required"})
				}

				if updatedListing.Price == 0 {
					errors.Errors = append(errors.Errors, responses.Error{FieldName: "price", Message: "Price is required"})
				}

				datePosted, err := time.Parse(time.RFC3339, updatedListing.DatePosted)
				if err == nil {
					updatedListing.DatePosted = datePosted.Format(time.RFC3339)
				}
				if datePosted.Year() == 1 {
					errors.Errors = append(errors.Errors, responses.Error{FieldName: "date_posted", Message: "Date Posted is required"})
				}

				if len(errors.Errors) > 0 {
					c.IndentedJSON(http.StatusBadRequest, errors)
					return
				}

				// Update the listing with the new values.
				listings[i] = updatedListing
				c.IndentedJSON(http.StatusOK, &responses.SuccessResponse{
					Data:    updatedListing,
					Status:  http.StatusOK,
					Message: "Listing was updated",
				})
				return
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, &responses.ErrorMessageResponse{
		Message: "listing not found",
		Error:   err.Error(),
		Status:  http.StatusNotFound,
	})
}

// @BasePath /listings/{id}
// PingExample godoc
// @Summary delete listing
// @Schemes
// @Description DeleteListing locates the listing whose ID value matches the id parameter sent by the client, then deletes it.
// @Tag listings
// @Param id path int true "Listing ID"
// @Accept json
// @Produce json
// @Success      200  {object}  responses.SuccessResponse
// @Failure      404  {object}  responses.ErrorMessageResponse
// @Router /listings/{id} [delete]
func (l *Listing) DeleteListing(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		// Loop over the slice of listings, looking for
		// a listing whose ID value matches the parameter.
		for i, l := range listings {
			if l.ID == id {
				// Remove the listing from the slice.
				listings = append(listings[:i], listings[i+1:]...)
				c.IndentedJSON(http.StatusOK, &responses.SuccessResponse{
					Message: "listing was deleted.",
					Data:    l,
					Status:  http.StatusOK,
				})
				return
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, &responses.ErrorMessageResponse{
		Message: "Listing not found",
		Error:   err.Error(),
		Status:  http.StatusNotFound,
	})
}

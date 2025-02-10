// Package routes Listing API.
//
//		Schemes: http
//		BasePath: /api
//		Version: 1.0.0
//	 Title: Classified Listings API
//		Host: localhost:8080
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
// swagger:meta
package routes

// swagger:route GET /api/listings listings GetListings
//
// GetListings returns all listings.
//
// Responses:
//
//	200: successResponse
func GetListings() {}

// swagger:route GET /api/listings/{id} listings GetListing
//
// GetListing returns one listing.
//
// Responses:
//
//	200: successResponse
//	404: errorMessageResponse
func GetListing() {}

// swagger:route POST /api/listings listings CreateListing
//
// CeeateListing creates a new listing.
//
// Responses:
//
//	200: successResponse
//	400: validationErrorResponse validation error response
func CreateListing() {}

// swagger:route PUT /api/listings listings UpdateListing
//
// UpdateListing updates a listing.
//
// Responses:
//
//		201: successResponse
//		400: errorMessageReponse error in decoding json
//		400: validationErrorResponse validation error response
//	    404: errorMessageReponse listing not found
func UpdateListing() {}

// swagger:route DELETE /api/listings listings DeleteListing
//
// DeleteListing deletes a listing.
//
// Responses:
//
//	200: successResponse
//	404: errorMessageResponse listing not found
func DeleteListing() {}

package main

import (
	"embed"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rossmoney/classified-listings-go/api/docs"
	"github.com/rossmoney/classified-listings-go/api/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// function to allow using env fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//go:embed static
var nuxtStatic embed.FS

// @title Classified Listings API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name Ross Money
// @contact.email ross.money@gmx.com
// @host localhost:8080
// @schemes http
// @BasePath /api
func main() {
	listing := &routes.Listing{}
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	//router.Use(static.Serve("/", static.EmbedFolder(nuxtStatic, "static")))
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	api := router.Group("/api")
	{
		api.GET("listings", listing.GetListings)
		api.GET("listings/:id", listing.GetListing)
		api.POST("listings", listing.CreateListing)
		api.PUT("listings/:id", listing.UpdateListing)
		api.DELETE("listings/:id", listing.DeleteListing)
	}
	router.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:" + getEnv("PORT", "8080"))
}

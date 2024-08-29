package main

import (
	"log"

	"code-challenge-backend/app"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var (
		r  = gin.Default()
		ds = app.NewDataStorage(viper.GetString("db"))
		h  = app.NewHandler(ds, viper.GetString("jwt_secret"))
		m  = app.NewMiddleware(viper.GetString("jwt_secret"))
	)

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	noAuth := r.Group("/api/v1")
	auth := r.Group("/api/v1")

	noAuth.POST("/register", h.Register)
	noAuth.POST("/login", h.Login)

	auth.Use(m.ValidateJWT)
	auth.GET("/seats", h.GetSeats)
	auth.POST("/book", h.BookSeat)

	log.Fatal(r.Run(":8080"))
}

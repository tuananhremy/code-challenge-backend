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
		r       = gin.Default()
		ds      = app.NewDataStorage(viper.GetString("db"))
		checkin = app.NewCheckInService(ds, viper.GetString("jwt_secret"))
	)
	go checkin.ReleaseBooking()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	noAuth := r.Group("/api/v1")
	noAuth.POST("/checkin", checkin.CheckIn)

	log.Fatal(r.Run(":8080"))
}

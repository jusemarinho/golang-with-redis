package main

import (
	"encoding/json"
	"fmt"
	"golang-with-redis/api/controllers"
	"golang-with-redis/config"
	"golang-with-redis/messaging"
	"golang-with-redis/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Error().Msg((err.Error()))
		return
	}

	var applicationConfig config.ApplicationConfig

	err = json.Unmarshal(file, &applicationConfig)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	rabbitMQ, err := messaging.NewRabbitMQ(&applicationConfig.AmqpConfig)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	pizzaService := services.NewPizzaService(rabbitMQ)

	pizzaController := controllers.NewPizzaController(pizzaService)

	r := gin.Default()

	r.POST("/pizzas", pizzaController.SavePizza)

	r.Run(fmt.Sprintf("%s:%d", applicationConfig.Application.Host, applicationConfig.Application.Port))
}

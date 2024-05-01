package services

import (
	"encoding/json"
	"golang-with-redis/messaging"

	"github.com/rs/zerolog/log"
)

type Pizza struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Price       float64      `json:"price"`
	Size        string       `json:"size"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type PizzaService struct {
	rabbitMQ *messaging.RabbitMQ
}

func NewPizzaService(rabbitMQ *messaging.RabbitMQ) *PizzaService {
	return &PizzaService{
		rabbitMQ: rabbitMQ,
	}
}

func (ps *PizzaService) CreatePizza(newPizza Pizza) {
	createdPizza := Pizza{
		ID:          "2",
		Name:        newPizza.Name,
		Price:       newPizza.Price,
		Size:        newPizza.Size,
		Ingredients: newPizza.Ingredients,
	}

	pizzaString, err := json.Marshal(createdPizza)
	if err != nil {
		log.Err(err)
		return
	}

	ps.rabbitMQ.PublishMessage(string(pizzaString))
}

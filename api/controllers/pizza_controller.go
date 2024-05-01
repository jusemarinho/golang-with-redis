package controllers

import (
	"golang-with-redis/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PizzaController struct {
	pizzaService *services.PizzaService
}

func NewPizzaController(ps *services.PizzaService) *PizzaController {
	return &PizzaController{
		pizzaService: ps,
	}
}

func (pz *PizzaController) SavePizza(c *gin.Context) {
	var reqBody services.Pizza 
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar JSON"})
		return
	}

	pz.pizzaService.CreatePizza(reqBody)

	c.Status(http.StatusAccepted)
}

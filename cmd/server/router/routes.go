package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	list   []domain.Ticket
}

func NewRouter(engine *gin.Engine, list []domain.Ticket) Router {
	return Router{
		engine,
		list,
	}
}

func (router *Router) MapRoutes() {
	repo := tickets.NewRepository(router.list)
	service := tickets.NewService(repo)
	ticketsHandler := handler.NewService(service)
	tickets := router.engine.Group("/ticket")
	{
		tickets.GET("/getByCountry/:dest", ticketsHandler.GetTicketsByCountry())
		tickets.GET("/getAverage/:dest", ticketsHandler.AverageDestination())
	}
	return
}

package tickets

import (
	"context"
	"errors"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(context.Context, string) ([]domain.Ticket, error)
	AverageDestination(context.Context, string) (float64, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetTotalTickets(c context.Context, destination string) (list []domain.Ticket, err error) {
	list, err = s.r.GetTicketByDestination(c, destination)

	return
}

func (s *service) AverageDestination(c context.Context, destination string) (total float64, err error) {
	list, err := s.r.GetTicketByDestination(c, destination)

	if err != nil {
		return
	}

	if len(list) == 0 {
		err = errors.New("no tickets found")
		return
	}

	for _, v := range list {
		total += v.Price
	}
	total = total / float64(len(list))

	return
}

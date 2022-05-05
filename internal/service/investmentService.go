package service

import (
	"context"
	"github.com/go-kit/log"
	"net/http"
	"test_robert_yofio/internal/entity"
	"test_robert_yofio/internal/repository"
)

type InvestmentService interface {
	CreditAssignment(context context.Context, data *interface{}) (entity.Response, error, int)
	Statistics(_ context.Context, data *interface{}) (entity.StatisticsResponse, error, int)
}

type investmentService struct {
	investmentRepository repository.InvestmentRepository
	logger               log.Logger
}

func NewInvestmentService(repo repository.InvestmentRepository, logger log.Logger) InvestmentService {
	return &investmentService{
		investmentRepository: repo,
		logger:               logger,
	}
}

func (i *investmentService) CreditAssignment(_ context.Context, data *interface{}) (entity.Response, error, int) {
	response, err := i.investmentRepository.CreditAssignment(data, i.logger)
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	return response, err, status
}

func (i *investmentService) Statistics(_ context.Context, _ *interface{}) (entity.StatisticsResponse, error, int) {
	response, err := i.investmentRepository.Statistics()
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	return response, err, status
}

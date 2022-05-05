package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"test_robert_yofio/internal/entity"
	"test_robert_yofio/internal/middleware"
	"test_robert_yofio/internal/service"
	"test_robert_yofio/internal/static"
)

type InvestmentEndpoints struct {
	CreditAssignmentEndpoint endpoint.Endpoint
}

func MakeInvestmentEndpoints(is *service.InvestmentService, im middleware.InvestmentMiddleware) InvestmentEndpoints {
	return InvestmentEndpoints{
		CreditAssignmentEndpoint: wrapEndpoint(makeCreditAssignmentEndpoint(*is), []endpoint.Middleware{im.AuthorizationInvestment()}),
	}
}

func makeCreditAssignmentEndpoint(is service.InvestmentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err, statusCode := is.CreditAssignment(ctx, &request)
		msg := static.MsgSuccessfully
		if err != nil {
			msg = err.Error()
		}
		investmentResponse := entity.InvestmentResponse{
			Data:       response,
			StatusCode: statusCode,
			Message:    msg,
		}
		return investmentResponse, err
	}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	if middlewares != nil {
		for _, m := range middlewares {
			e = m(e)
		}
	}
	return e
}

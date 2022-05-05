package middleware

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"regexp"
	"strconv"
	"test_robert_yofio/internal/entity"
	"test_robert_yofio/internal/static"
)

var (
	ErrUnauthorizatedInvestment = errors.New(static.MsgUnauthorizatedInvestment)
	ErrInvalidInvestment        = errors.New(static.MsgInvalidInvestment)
	IsMultiple100               = regexp.MustCompile(`.00$`).MatchString
)

type investmentMiddleware struct {
	logger log.Logger
}

type (
	InvestmentMiddleware interface {
		AuthorizationInvestment() endpoint.Middleware
	}
)

func NewInvestmentMiddleware(logger log.Logger) InvestmentMiddleware {
	return &investmentMiddleware{
		logger: logger,
	}
}

func (im *investmentMiddleware) AuthorizationInvestment() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (response interface{}, err error) {
			traces := []interface{}{}
			err = errors.New(static.ValueEmpty)
			investment := req.(entity.InvestmentRequest).Investment
			if !IsMultiple100(strconv.Itoa(investment)) {
				traces = []interface{}{static.KeyType, static.ERROR, static.KeyMessage, ErrUnauthorizatedInvestment.Error()}
				err = ErrUnauthorizatedInvestment
			}
			if len(traces) > 0 {
				LoggingOperation(im.logger, traces...)
				return nil, err
			}
			return next(ctx, req)
		}
	}
}

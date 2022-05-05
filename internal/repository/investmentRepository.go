package repository

import (
	"errors"
	"github.com/go-kit/log"
	"gopkg.in/mgo.v2/bson"
	"test_robert_yofio/internal/entity"
	"test_robert_yofio/internal/function"
	"test_robert_yofio/internal/middleware"
	"test_robert_yofio/internal/static"
)

var (
	ErrResponseObjectExists       = errors.New(static.MsgResponseObjectExists)
	ErrResponseServerErrorNoID    = errors.New(static.MsgResponseServerErrorNoID)
	ErrResponseServerErrorWrongID = errors.New(static.MsgResponseServerErrorWrongID)
	ErrResponseServerErrorNoData  = errors.New(static.MsgResponseServerErrorNoData)
)

type InvestmentRepository interface {
	CreditAssignment(r *interface{}, logger log.Logger) (entity.Response, error)
}

type investmentRepository struct {
	creditAssignment creditAssignment.CreditAssigner
}

func NewInvestmentRepository(creditAssignment creditAssignment.CreditAssigner) InvestmentRepository {
	return &investmentRepository{
		creditAssignment: creditAssignment,
	}
}

//Creating Person
func (ir *investmentRepository) CreditAssignment(r *interface{}, logger log.Logger) (entity.Response, error) {
	traces := []interface{}{}
	object := ir.ToEntityObject(*r)
	credit300, credit500, credit700, err := ir.creditAssignment.Assign(int32(object.Investment))
	response := entity.Response{}
	if err != nil {
		traces = []interface{}{static.KeyType, static.ERROR, static.KeyMessage, static.MsgErrorOperation}
		middleware.LoggingOperation(logger, traces...)
		return response, err
	}
	response.CreditType300, response.CreditType500, response.CreditType700 = credit300, credit500, credit700
	traces = []interface{}{static.KeyType, static.SUCCESS, static.KeyMessage, static.MsgSuccessfully}
	middleware.LoggingOperation(logger, traces...)
	return response, nil

}

//FUNCIONES AUXILIARES

func (ir *investmentRepository) ToEntityObject(i interface{}) entity.InvestmentRequest {
	obj := entity.InvestmentRequest{}
	bsonBytes, _ := bson.Marshal(i)
	_ = bson.Unmarshal(bsonBytes, &obj)
	return obj
}

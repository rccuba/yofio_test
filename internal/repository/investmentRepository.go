package repository

import (
	"errors"
	"github.com/go-kit/log"
	"gopkg.in/mgo.v2/bson"
	"math"
	"test_robert_yofio/internal/db"
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
	Statistics() (entity.StatisticsResponse, error)
}

type investmentRepository struct {
	creditAssignment creditAssignment.CreditAssigner
	connMongoDb      *db.MongoConnection
}

func NewInvestmentRepository(creditAssignment creditAssignment.CreditAssigner, connMongoDb *db.MongoConnection) InvestmentRepository {
	return &investmentRepository{
		creditAssignment: creditAssignment,
		connMongoDb:      connMongoDb,
	}
}

//Credit Assignment
func (ir *investmentRepository) CreditAssignment(r *interface{}, logger log.Logger) (entity.Response, error) {
	traces := []interface{}{}
	object := ir.ToEntityObject(*r)
	credit300, credit500, credit700, err := ir.creditAssignment.Assign(int32(object.Investment))
	response := entity.Response{}
	if err != nil {
		traces = []interface{}{static.KeyType, static.ERROR, static.KeyMessage, static.MsgErrorOperation}
	} else {
		response.CreditType300, response.CreditType500, response.CreditType700 = credit300, credit500, credit700
		traces = []interface{}{static.KeyType, static.SUCCESS, static.KeyMessage, static.MsgSuccessfully}
	}
	middleware.LoggingOperation(logger, traces...)
	ir.connMongoDb.InsertData(static.CollectionCreditAssignment, entity.StatisticsRequest{
		CreditAssignment: response,
		Investment:       int32(object.Investment),
	})
	return response, err
}

//Statistics
func (ir *investmentRepository) Statistics() (entity.StatisticsResponse, error) {
	items, err := ir.connMongoDb.GetFindData(static.CollectionCreditAssignment, bson.M{}, bson.M{}, static.FieldInvestment, static.SortAsc)
	statistics := make([]entity.StatisticsRequest, 0)
	for _, item := range items {
		elem := &entity.StatisticsRequest{}
		bsonBytes, _ := bson.Marshal(item)
		_ = bson.Unmarshal(bsonBytes, &elem)
		statistics = append(statistics, *elem)
	}
	countSuccess, countError := 0, 0
	investmentSuccess, investmentError := int32(0), int32(0)
	for _, item := range statistics {
		if item.CreditAssignment.CreditType300+item.CreditAssignment.CreditType500+item.CreditAssignment.CreditType700 > 0 {
			countSuccess += 1
			investmentSuccess = investmentSuccess + item.Investment
		} else {
			countError += 1
			investmentError = investmentError + item.Investment
		}
	}
	return entity.StatisticsResponse{
		AssignmentInTotal:        int32(countSuccess + countError),
		AssignmentSuccess:        int32(countSuccess),
		AssignmentError:          int32(countError),
		InvestmentAverageSuccess: math.Round(float64(investmentSuccess) / float64(countSuccess)),
		InvestmentAverageError:   math.Round(float64(investmentError) / float64(countError)),
	}, err
}

//FUNCIONES AUXILIARES
func (ir *investmentRepository) ToEntityObject(i interface{}) entity.InvestmentRequest {
	obj := entity.InvestmentRequest{}
	bsonBytes, _ := bson.Marshal(i)
	_ = bson.Unmarshal(bsonBytes, &obj)
	return obj
}

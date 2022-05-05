package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"test_robert_yofio/internal/app"
	"test_robert_yofio/internal/config"
	"test_robert_yofio/internal/entity"
	"test_robert_yofio/internal/static"
	"testing"
)

var a app.App

func TestMain(m *testing.M) {
	if err, isConfigurable := config.ConfigEnv(); !isConfigurable {
		fmt.Printf(""+static.MsgResponseStartError+", %s", err)
	} else {
		a = App()
		code := m.Run()
		os.Exit(code)
	}
}

func App() app.App {
	a := app.App{}
	_ = a.Initialize()
	return a
}

func TestCreditAssignment(t *testing.T) {
	t.Run(static.MsgTestUnsupportedHTTPMethod, func(t *testing.T) {
		payload := []byte(`{"investment" : 1}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssignmentTest(http.MethodOptions, url, bytes.NewBuffer(payload), static.ValueEmpty, static.ValueEmpty)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 6000", func(t *testing.T) {
		payload := []byte(`{"investment" : 6000}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssignmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 300", func(t *testing.T) {
		payload := []byte(`{"investment" : 300}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssignmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 500", func(t *testing.T) {
		payload := []byte(`{"investment" : 500}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssignmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 700", func(t *testing.T) {
		payload := []byte(`{"investment" : 700}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssignmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestUndeliveredAmount, func(t *testing.T) {
		payload := []byte(`{"investment" : 200}`)
		msg := static.MsgUndeliveredAmount + " : 200"
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssignmentTest(http.MethodPost, url, bytes.NewBuffer(payload), msg, msg)
		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestUnauthorizatedInvestment, func(t *testing.T) {
		payload := []byte(`{"investment" : 5005}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssignmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgUnauthorizatedInvestment, static.MsgUnauthorizatedInvestment)
		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
}

func TestStatistics(t *testing.T) {
	t.Run(static.MsgTestStatistics, func(t *testing.T) {
		url := static.URLStatistics
		response := ResponseStatisticsTest(http.MethodGet, url, bytes.NewBuffer(nil), static.ValueEmpty, static.ValueEmpty)
		assert.GreaterOrEqual(t, response.AssignmentInTotal, int32(0))
		assert.GreaterOrEqual(t, response.AssignmentSuccess, int32(0))
		assert.GreaterOrEqual(t, response.AssignmentError, int32(0))
		assert.GreaterOrEqual(t, response.InvestmentAverageSuccess, float64(0))
		assert.GreaterOrEqual(t, response.InvestmentAverageError, float64(0))
	})
}

//FUNCIONES EXTRAS
func ResponseCreditAssignmentTest(method string, url string, body *bytes.Buffer, eval string, defaultMessage string) (string, string) {
	request, _ := http.NewRequest(method, url, body)
	fmt.Println(request)
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	responseBody, err := ResponseToJSON(response.Body.String())
	fmt.Println(responseBody)
	return func() (string, string) {
		if err != nil {
			return response.Body.String(), eval
		}
		return responseBody[static.KeyMessage].(string), defaultMessage
	}()
}

func ResponseStatisticsTest(method string, url string, body *bytes.Buffer, eval string, defaultMessage string) entity.StatisticsResponse {
	request, _ := http.NewRequest(method, url, body)
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	responseBody, _ := ResponseToJSON(response.Body.String())
	jsonStr, _ := json.Marshal(responseBody)
	var elem entity.StatisticsResponse
	json.Unmarshal(jsonStr, &elem)
	return elem
}

func ResponseToJSON(responseBody string) (map[string]interface{}, error) {
	var JSON map[string]interface{}
	err := json.Unmarshal([]byte(responseBody), &JSON)
	return JSON, err
}

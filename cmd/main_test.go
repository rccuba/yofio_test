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

func ResponseToJSON(responseBody string) (map[string]interface{}, error) {
	var JSON map[string]interface{}
	err := json.Unmarshal([]byte(responseBody), &JSON)
	return JSON, err
}

func TestCreditAssignment(t *testing.T) {
	t.Run(static.MsgTestUnsupportedHTTPMethod, func(t *testing.T) {
		payload := []byte(`{"investment" : 1}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssigmentTest(http.MethodOptions, url, bytes.NewBuffer(payload), static.ValueEmpty, static.ValueEmpty)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 6000", func(t *testing.T) {
		payload := []byte(`{"investment" : 6000}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssigmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 300", func(t *testing.T) {
		payload := []byte(`{"investment" : 300}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssigmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 500", func(t *testing.T) {
		payload := []byte(`{"investment" : 500}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssigmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestValidValueInvestment+" : 700", func(t *testing.T) {
		payload := []byte(`{"investment" : 700}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssigmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgSuccessfully, static.MsgSuccessfully)

		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestUndeliveredAmount, func(t *testing.T) {
		payload := []byte(`{"investment" : 200}`)
		msg := static.MsgUndeliveredAmount + " : 200"
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssigmentTest(http.MethodPost, url, bytes.NewBuffer(payload), msg, msg)
		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
	t.Run(static.MsgTestUnauthorizatedInvestment, func(t *testing.T) {
		payload := []byte(`{"investment" : 5005}`)
		url := static.URLCreditAssignment
		message, eval := ResponseCreditAssigmentTest(http.MethodPost, url, bytes.NewBuffer(payload), static.MsgUnauthorizatedInvestment, static.MsgUnauthorizatedInvestment)
		assert.Equal(t, message, eval, static.MsgTestEXPECTED+" "+message)
	})
}

//FUNCIONES EXTRAS
func ResponseCreditAssigmentTest(method string, url string, body *bytes.Buffer, eval string, defaultMessage string) (string, string) {
	request, _ := http.NewRequest(method, url, body)
	response := httptest.NewRecorder()
	a.Router.ServeHTTP(response, request)
	responseBody, err := ResponseToJSON(response.Body.String())
	return func() (string, string) {
		if err != nil {
			return response.Body.String(), eval
		}
		return responseBody[static.KeyMessage].(string), defaultMessage
	}()
}

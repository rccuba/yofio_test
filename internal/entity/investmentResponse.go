package entity

type InvestmentResponse struct {
	StatusCode int      `json:"statusCode"`
	Message    string   `json:"message"`
	Data       Response `json:"data"`
}

type Response struct {
	CreditType300 int32 `json:"“credit_type_300”"`
	CreditType500 int32 `json:"“credit_type_500”"`
	CreditType700 int32 `json:"“credit_type_700”"`
}

package entity

type InvestmentRequest struct {
	Investment int `json:"investment,omitempty" bson:"investment,omitempty"`
}

type GenericRequest struct {
}

package entity

type StatisticsRequest struct {
	CreditAssignment Response `json:"“credit_assignment,omitempty" bson:"credit_assignment,omitempty"`
	Investment       int32    `json:"“investment,omitempty" bson:"investment,omitempty"`
}

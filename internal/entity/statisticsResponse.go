package entity

type StatisticsResponse struct {
	AssignmentInTotal        int32   `json:"“total_asignaciones_realizadas”"`
	AssignmentSuccess        int32   `json:"“total_asignaciones_exitosas”"`
	AssignmentError          int32   `json:"“total_asignaciones_no_exitosas”"`
	InvestmentAverageSuccess float64 `json:"“promedio_inversion_exitosa”"`
	InvestmentAverageError   float64 `json:"“promedio_inversion_no_exitosa”"`
}

package app

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"test_robert_yofio/internal/coders"
)

// CreditAssignment godoc
// @Summary Asignación de créditos
// @Description Comienza a asignar créditos a partir de un monto y retorna una de las posibles opciones
// @Accept  plain
// @Produce  json
// @Param parameters body entity.InvestmentRequest true "PAYLOAD DE ENTRADA (VALOR ENTERO)"
// @Success 200 {object} entity.InterfaceAPI
// @Router /credit-assignment [post]
func (a *App) CreditAssignment(options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		a.InvestmentEndpoints.CreditAssignmentEndpoint,
		coders.DecodeInvestmentRequest,
		coders.RespondWithJSON,
		options...,
	)
}

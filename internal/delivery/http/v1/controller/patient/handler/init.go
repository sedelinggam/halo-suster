package patientHandler

import (
	patientService "halo-suster/internal/service/patient"

	"github.com/go-playground/validator/v10"
)

type patientHandler struct {
	patientService patientService.PatientService
	val            *validator.Validate
}

func NewHandler(patientService patientService.PatientService, val *validator.Validate) *patientHandler {
	return &patientHandler{
		patientService: patientService,
		val:            val,
	}
}

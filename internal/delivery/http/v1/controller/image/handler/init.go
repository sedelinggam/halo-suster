package imageHandler

import (
	"github.com/go-playground/validator/v10"
)

type imageHandler struct {
	val *validator.Validate
}

func NewHandler(val *validator.Validate) *imageHandler {
	return &imageHandler{
		val: val,
	}
}

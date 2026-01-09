package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationEror(err error) []string{
	var errorMessages []string

	if validationErrors,ok := err.(validator.ValidationErrors); ok  {
		for _,e := range validationErrors{
			switch e.Tag(){
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s wajib di isi",e.Field()))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf("%s harus berupa format Email yang valid",e.Field()))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s minimal %s karakter",e.Field(),e.Param()))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("%s tidak valid",e.Field()))
			}
		}
	}else{
		errorMessages = append(errorMessages, "Format INput tidak valid")
	}
	return errorMessages
}
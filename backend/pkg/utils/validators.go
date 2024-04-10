package utils

import (
	"fmt"
	"reflect"
	"regexp"
)

func ValidateStruct(s interface{}) (err error) {
	structType := reflect.TypeOf(s)

	structVal := reflect.ValueOf(s)
	fieldNum := structVal.NumField()

	var errors string

	for i := 0; i < fieldNum; i++ {

		field := structVal.Field(i)
		fieldName := structType.Field(i).Name

		isSet := field.IsValid() && !field.IsZero()

		if !isSet {
			errors += fmt.Sprintf("Falta el campo %s ", fieldName)
		}

	}

	if errors == "" {
		return nil
	}

	return fmt.Errorf(errors)
}

func ValidateEmail(email string) error {
	match, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if !match {
		return fmt.Errorf("El correo no es valido")
	}
	return nil
}

func ValidatePhone(phone string) error {
	match, _ := regexp.MatchString(`^5[0-9]{9}$`, phone)
	if !match {
		return fmt.Errorf("El telefono no es valido")
	}
	return nil
}

func ValidatePassword(password string) error {

	if len(password) > 12 || len(password) < 6 {
		return fmt.Errorf("La contrasena no es valida")
	}

	match, _ := regexp.MatchString(`^.*[a-z].*[A-Z].*[@$&].*\d.*|.*[a-z].*[A-Z].*\d.*[@$&].*$|.*[A-Z].*[a-z].*[@$&].*\d.*|.*[A-Z].*[a-z].*\d.*[@$&].*$|.*[@$&].*[a-z].*[A-Z].*\d.*|.*[@$&].*[a-z].*\d.*[A-Z].*$|.*\d.*[@$&].*[a-z].*[A-Z].*$|.*\d.*[@$&].*[A-Z].*[a-z].*$|.*[@$&].*\d.*[a-z].*[A-Z].*$|.*[@$&].*\d.*[A-Z].*[a-z].*$`, password)
	if !match {
		return fmt.Errorf("La contrasena no es valida")
	}

	return nil
}

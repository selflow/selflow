package config

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"text/template"
	"time"
)

const (
	durationValidatorKey = "duration"
	templateKey          = "template"
	identifierKey        = "identifier"
	identifierKeyKey     = "identifierKey"
)

func validateGoTemplate(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind().String() != "string" {
		return false
	}
	_, err := template.New("").Parse(fl.Field().String())

	println(err)
	return err == nil
}

const IdentifierRegex = "^[a-zA-Z][a-zA-Z0-9_\\-]*$"

func isValidIdentifier(identifier string) bool {
	return regexp.MustCompile(IdentifierRegex).MatchString(identifier)
}

func validateIdentifier(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind().String() != "string" {
		return false
	}
	return isValidIdentifier(field.String())
}

func validateIdentifierMap(fl validator.FieldLevel) bool {
	keys := fl.Field().MapKeys()
	for _, key := range keys {
		if !isValidIdentifier(key.String()) {
			return false
		}
	}

	return true
}

func validateDuration(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind().String() != "string" {
		return false
	}

	_, err := time.ParseDuration(field.String())

	return err == nil
}

func InitValidation() *validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation(templateKey, validateGoTemplate)
	if err != nil {
		panic(err)
	}
	err = validate.RegisterValidation(identifierKey, validateIdentifier)
	if err != nil {
		panic(err)
	}
	err = validate.RegisterValidation(identifierKeyKey, validateIdentifierMap)
	if err != nil {
		panic(err)
	}
	err = validate.RegisterValidation(durationValidatorKey, validateDuration)
	if err != nil {
		panic(err)
	}

	return validate
}

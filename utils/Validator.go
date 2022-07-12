package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en2 "gopkg.in/go-playground/validator.v9/translations/en"
)

func Validate(req interface{}) string {
	errorString := ""
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en2.RegisterDefaultTranslations(v, trans)

	if err := v.Struct(req); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			errorString += fmt.Sprintf(v.Translate(trans)) + ", "
		}
	}
	if errorString == "" {
		return errorString
	}
	return errorString[:len(errorString)-2]
}

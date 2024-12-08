package translator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
)

type Translator struct {
	trans ut.Translator
}

func NewTranslator(validate *validator.Validate) *Translator {
	t := Translator{}
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en2.RegisterDefaultTranslations(validate, trans)
	t.trans = trans

	return &t
}

func (t *Translator) TranslateError(err error) map[string]string {
	var (
		validatorErrs validator.ValidationErrors
		mappedErrors  = map[string]string{}
	)

	if err == nil {
		return nil
	}
	errors.As(err, &validatorErrs)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(t.trans)
		mappedErrors[e.StructField()+"."+e.Tag()] = translatedErr
	}

	return mappedErrors
}

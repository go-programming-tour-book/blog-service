package validator

import (
	"reflect"
	"sync"

	"gopkg.in/go-playground/validator.v9"
)

type ValidatorV9 struct {
	Once     sync.Once
	Validate *validator.Validate
}

func NewValidatorV9() *ValidatorV9 {
	return &ValidatorV9{}
}

func (v *ValidatorV9) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.Validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *ValidatorV9) Engine() interface{} {
	v.lazyinit()
	return v.Validate
}

func (v *ValidatorV9) lazyinit() {
	v.Once.Do(func() {
		v.Validate = validator.New()
		v.Validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

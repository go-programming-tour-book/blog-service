package validator

import (
	"reflect"
	"sync"

	val "github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Once     sync.Once
	Validate *val.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

func (v *CustomValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.Validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *CustomValidator) Engine() interface{} {
	v.lazyinit()
	return v.Validate
}

func (v *CustomValidator) lazyinit() {
	v.Once.Do(func() {
		v.Validate = val.New()
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

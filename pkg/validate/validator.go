package validate

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

var _ binding.StructValidator = (*Validator)(nil)

func NewValidator() (*Validator, error) {
	v := &Validator{
		validate: validator.New(),
	}

	v.validate.SetTagName("binding")
	v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("label")
	})

	// 使用中文进行交互
	uni := ut.New(zh.New())
	v.trans, _ = uni.GetTranslator("zh")

	err := zh_translations.RegisterDefaultTranslations(v.validate, v.trans)
	if err != nil {
		return nil, err
	}

	// 注册自定义验证规则

	for key, f := range Rules {
		if err = v.validate.RegisterValidation(key, f); err != nil {
			return nil, err
		}
	}

	// 注册自定义错误内容
	for key, val := range RulesMsg {
		if err = v.validate.RegisterTranslation(key, v.trans, val.RegisterTranslationsFunc, val.TranslationFunc); err != nil {
			return nil, err
		}
	}

	return v, nil
}

// ValidateStruct 接收任何类型的类型，但只执行 struct 或指向 struct 类型的指针。
func (v *Validator) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		if value.Elem().Kind() != reflect.Struct {
			return v.ValidateStruct(value.Elem().Interface())
		}

		return v.validateStruct(obj)
	case reflect.Struct:
		return v.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(binding.SliceValidationError, 0)

		for i := range count {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}

		if len(validateRet) == 0 {
			return nil
		}

		return validateRet
	default:
		return nil
	}
}

// validateStruct receives struct type.
func (v *Validator) validateStruct(obj any) error {
	err := v.validate.Struct(obj)
	if err == nil {
		return nil
	}

	var verrs validator.ValidationErrors
	if !errors.As(err, &verrs) {
		return err
	}

	msg := make([]string, 0)
	for _, val := range verrs.Translate(v.trans) {
		msg = append(msg, val)
	}

	return errors.New(strings.Join(msg, ","))
}

func (v *Validator) Engine() any {
	return v.validate
}

package validator

import (
	"fmt"
	"gin_demo/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

//用户权限 验证

func Validator(data interface{}) (string, int) {
	validate := validator.New() // 创建验证规则

	uni := ut.New(zh_Hans_CN.New())

	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err", err)
	}

	//映射label标签
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}

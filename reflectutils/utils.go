package reflectutils

import (
	"reflect"
)

// NewInstance 反射实例化
func NewInstance(fillType interface{}) interface{} {
	targetType := reflect.TypeOf(fillType)
	targetInstance := reflect.New(targetType).Interface()
	return targetInstance
}

// SetProperty 设置指定属性
func SetProperty(instance interface{}, fieldName string, val interface{}) {
	targetVal := reflect.ValueOf(instance)
	field := targetVal.FieldByName(fieldName)
	if field.CanSet() {
		field.Set(reflect.ValueOf(val))
	}
}

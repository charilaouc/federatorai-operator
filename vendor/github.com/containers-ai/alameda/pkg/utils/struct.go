package utils

import (
	"fmt"
	"reflect"
)

func ExtractField(entities interface{}, fieldName string) reflect.StructField {
	var structField reflect.StructField

	fieldType := reflect.TypeOf(entities).Elem()
	switch fieldType.Kind() {
	case reflect.Struct:
		structField, _ = fieldType.FieldByName(fieldName)
	case reflect.Slice:
		structField, _ = fieldType.Elem().FieldByName(fieldName)
	default:
		fmt.Printf("error: data type(%d) of entity is not supported", fieldType.Kind())
	}

	return structField
}

package entities

import (
	"github.com/containers-ai/alameda/pkg/utils"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/golang/protobuf/ptypes/timestamp"
	"reflect"
	"strconv"
	"time"
)

var (
	scope            = log.RegisterScope("entity", "datahub entity library", 0)
	ColumnStartIndex = 2
)

type Row struct {
	Time    *timestamp.Timestamp
	Columns []string
	Values  []string
}

type Entity interface {
	Populate(entity interface{}, timestamp *timestamp.Timestamp, columns, values []string)
	Row(entity interface{}, fields []string) *Row
}

type Measurement struct {
}

type DatahubEntity struct {
}

func (p *DatahubEntity) Populate(entity interface{}, timestamp *timestamp.Timestamp, columns, values []string) {
	fields := reflect.TypeOf(entity).Elem()

	// Populate Time filed
	ts := utils.Timestamp(timestamp)
	timeField := reflect.ValueOf(entity).Elem().FieldByName("Time")
	timeField.Set(reflect.ValueOf(ts))

	// Populate other tags and fields
	for index, column := range columns {
		// Index is started at 2 to skip DatahubEntity and Measurement fields
		for i := ColumnStartIndex; i < fields.NumField(); i++ {
			if column == fields.Field(i).Tag.Get("json") {
				fieldValue := reflect.ValueOf(entity).Elem().Field(i)
				switch fieldValue.Kind() {
				case reflect.Bool:
					valueBool, _ := strconv.ParseBool(values[index])
					fieldValue.SetBool(valueBool)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
					valueInt, _ := strconv.ParseInt(values[index], 10, 32)
					fieldValue.SetInt(valueInt)
				case reflect.Int64:
					valueInt, _ := strconv.ParseInt(values[index], 10, 64)
					fieldValue.SetInt(valueInt)
				case reflect.Float32:
					valueFloat, _ := strconv.ParseFloat(values[index], 32)
					fieldValue.SetFloat(valueFloat)
				case reflect.Float64:
					valueFloat, _ := strconv.ParseFloat(values[index], 64)
					fieldValue.SetFloat(valueFloat)
				case reflect.String:
					fieldValue.SetString(values[index])
				default:
					scope.Errorf("field type(%s) not supported", fieldValue.Kind().String())
				}
				break
			}
		}
	}
}

func (p *DatahubEntity) Row(entity interface{}, fields []string) *Row {
	row := Row{Time: utils.TimestampProto(reflect.ValueOf(entity).Elem().FieldByName("Time").Interface().(*time.Time))}

	// If fields is empty which means to iterate through all the fields of the entity
	values := reflect.TypeOf(entity).Elem()
	fieldNames := fields
	if len(fieldNames) == 0 {
		// Index is started at 2 to skip DatahubEntity and Measurement fields
		for i := ColumnStartIndex; i < values.NumField(); i++ {
			fieldNames = append(fieldNames, values.Field(i).Name)
		}
	}

	for _, name := range fieldNames {
		// Index is started at 2 to skip DatahubEntity and Measurement fields
		for i := ColumnStartIndex; i < values.NumField(); i++ {
			if name == values.Field(i).Name {
				fieldValue := reflect.ValueOf(entity).Elem().Field(i)
				switch fieldValue.Kind() {
				case reflect.Bool:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatBool(fieldValue.Bool()))
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatInt(fieldValue.Int(), 10))
				case reflect.Float32:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 32))
				case reflect.Float64:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64))
				case reflect.String:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, fieldValue.String())
				default:
					scope.Errorf("field type(%s) not supported", fieldValue.Kind().String())
				}
				break
			}
		}
	}

	return &row
}

func (p *DatahubEntity) Tags(entity interface{}) []string {
	tags := make([]string, 0)

	fieldType := reflect.TypeOf(entity)
	// Index is started at 3 to skip DatahubEntity, Measurement and Time fields
	for i := ColumnStartIndex + 1; i < fieldType.NumField(); i++ {
		if fieldType.Field(i).Tag.Get("column") == "tag" {
			tags = append(tags, fieldType.Field(i).Tag.Get("json"))
		}
	}

	return tags
}

func (p *DatahubEntity) Fields(entity interface{}) []string {
	fields := make([]string, 0)

	fieldType := reflect.TypeOf(entity)
	// Index is started at 3 to skip DatahubEntity, Measurement and Time fields
	for i := ColumnStartIndex + 1; i < fieldType.NumField(); i++ {
		if fieldType.Field(i).Tag.Get("column") == "field" {
			fields = append(fields, fieldType.Field(i).Tag.Get("json"))
		}
	}

	return fields
}

func (p *DatahubEntity) TagNames(entity interface{}) []string {
	tagNames := make([]string, 0)

	fieldType := reflect.TypeOf(entity)
	// Index is started at 3 to skip DatahubEntity, Measurement and Time fields
	for i := ColumnStartIndex + 1; i < fieldType.NumField(); i++ {
		if fieldType.Field(i).Tag.Get("column") == "tag" {
			tagNames = append(tagNames, fieldType.Field(i).Name)
		}
	}

	return tagNames
}

func (p *DatahubEntity) FieldNames(entity interface{}) []string {
	fields := make([]string, 0)

	fieldType := reflect.TypeOf(entity)
	// Index is started at 3 to skip DatahubEntity, Measurement and Time fields
	for i := ColumnStartIndex + 1; i < fieldType.NumField(); i++ {
		if fieldType.Field(i).Tag.Get("column") == "field" {
			fields = append(fields, fieldType.Field(i).Name)
		}
	}

	return fields
}

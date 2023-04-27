package simple

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func StructToMap(obj interface{}, excludes ...string) map[string]interface{} {
	var data = make(map[string]interface{})
	keys := reflect.TypeOf(obj)
	values := reflect.ValueOf(obj)
	fillMap(data, keys, values, excludes...)
	return data
}

func fillMap(data map[string]interface{}, keys reflect.Type, values reflect.Value, excludes ...string) {
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	if keys.Kind() == reflect.Ptr {
		keys = keys.Elem()
	}

	for i := 0; i < keys.NumField(); i++ {
		keyField := keys.Field(i)
		valueField := values.Field(i)

		if keyField.Anonymous {
			fillMap(data, keyField.Type, valueField, excludes...)
		} else {
			if !ContainsIgnoreCase(keyField.Name, excludes) {
				jsonTag := keyField.Tag.Get("json")
				if len(jsonTag) > 0 {
					data[jsonTag] = valueField.Interface()
				} else {
					data[keyField.Name] = valueField.Interface()
				}
			}
		}
	}
}

func MapToStruct(obj interface{}, data map[string]interface{}) error {
	for k, v := range data {
		err := setField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj ", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value ", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type ")
	}
	structFieldValue.Set(val)
	return nil
}

func StructFields(s interface{}) []reflect.StructField {
	t := StructTypeOf(s)
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}

	var results []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		results = append(results, f)
	}
	return results
}

func StructName(s interface{}) string {
	t := StructTypeOf(s)
	return t.Name()
}

func StructTypeOf(s interface{}) reflect.Type {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func InterfaceToString(value interface{}) string {

	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
func InterfaceToStruct(cacheData string, transferred interface{}) error {
	if len(cacheData) <= 0 {
		return errors.New("非法字符串")
	}
	newData := InterfaceToString(cacheData)

	if err := json.Unmarshal([]byte(newData), transferred); err != nil {
		return errors.New("无法转化")
	}

	return nil
}

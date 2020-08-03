package util

import (
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"

	"github.com/siddontang/go/bson"
)

func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case string:
		return fmt.Sprint(inter.(string))
		break
	case int:
		return fmt.Sprint(inter.(int))
		break
	case int64:
		return fmt.Sprint(inter.(int64))
		break
	case float64:
		return fmt.Sprint(inter.(float64))
		break
	}
	return ""
}

//判断interface是否为nil
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

//struct转成map
func StructToMap(st interface{}) map[string]interface{} {
	types := reflect.TypeOf(st)
	values := reflect.ValueOf(st)

	var data = make(map[string]interface{})
	for i := 0; i < types.NumField(); i++ {
		k := types.Field(i).Name
		var v = values.Field(i).Interface()
		if !IsNil(v) {
			switch v.(type) {
			case int:
				data[k] = v
			case *string:
				vv := v.(*string)
				data[k] = *vv
			case string:
				data[k] = v
			}
		}
	}

	return data
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	} else if valueMap, ok := value.(bson.M); ok {
		newMap := make(bson.M)
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
	}
	return value
}

//严格的struct拷贝
func StrictStructCopy(src interface{}, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}

//利用json做中间值进行拷贝
func StructCopyUseJson(src interface{}, dst interface{}) {
	//srcJson, _ := json.Marshal(src)
	//_ = json.Unmarshal(srcJson, dst)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonBytes, _ := json.Marshal(src)
	_ = json.Unmarshal(jsonBytes, dst)
}

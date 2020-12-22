package mzjstruct

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func CopyStruct(src, dst interface{}) {
	bt, _ := json.Marshal(src)
	json.Unmarshal(bt, dst)
}

func CopyStructs(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name
		t := value.Kind().String()
		fmt.Println(t)
		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		switch value.Kind() {
		case reflect.Struct:
			CopyStructs(&value, &dvalue)
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
			// 递归查询元素类型
			CopyStructs(&value, &dvalue)
		case reflect.String, reflect.Int, reflect.Int64, reflect.Int32, reflect.Bool, reflect.Int16, reflect.Int8:
			dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下
		default:
			dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
		}
	}
}
func RemoveSliceMap(a []interface{}) (ret []interface{}) {
	n := len(a)
	for i := 0; i < n; i++ {
		state := false
		for j := i + 1; j < n; j++ {
			if j > 0 && reflect.DeepEqual(a[i], a[j]) {
				state = true
				break
			}
		}
		if !state {
			ret = append(ret, a[i])
		}
	}
	return
}

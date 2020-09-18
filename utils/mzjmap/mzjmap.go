package mzjmap

import (
	"fmt"
	"reflect"
)

//EntityToEntity 实体转实体
func EntityToEntity(rqs interface{}, resp interface{}) error {
	v := reflect.ValueOf(rqs)
	t := v.Type()
	k := t.Kind()
	v2 := reflect.ValueOf(resp)
	t2 := v2.Type()
	k2 := t2.Kind()
	switch k {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			//field := v.Field(i)
			//t2.Field(i).Name = t.Field(i).Name
			switch k2 {
			case reflect.Struct:
				for j := 0; j < v2.NumField(); j++ {
					//t2.Field(i).Name = t.Field(i).Name
					if t.Field(i).Name == t2.Field(j).Name {
						//v2.Field(j) = field //Interface
					}

				}
			}
		}
	}
	return nil
}
func MapToInterface(req interface{},resp ...interface{})error  {
	return nil
}
func main()  {
	type Test struct {
		Name string
	}
	ts:=[]Test{}
	//mp :=[]map[string]string{}
	v:=reflect.ValueOf(ts)
	t:=v.Type()
	fmt.Println(t.Kind())
	switch t.Kind() {
	case reflect.Slice:

	}
}
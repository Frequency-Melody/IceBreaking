package util

import (
	"IceBreaking/log"
	"IceBreaking/response"
	"reflect"
)

// StructMapByFieldName 把 model 转成 dto
// 代码来自 https://github.com/Codexiaoyi/go-mapper
func ModelToDto(src, dest interface{}) interface{} {
	//fmt.Println(reflect.TypeOf(src).Kind())
	//fmt.Println(reflect.TypeOf(dest).Kind())
	if reflect.TypeOf(src).Kind() != reflect.Ptr || reflect.TypeOf(dest).Kind() != reflect.Ptr {
		//return errors.New("src and dst must be addressable")
		log.Sugar().Error("dto反射调用失败")
		return response.ServerUnknownError
	}

	dic := make(map[string]reflect.Value)
	srcPtr := reflect.ValueOf(src).Elem()
	destPtr := reflect.ValueOf(dest).Elem()

	//存储src字段信息
	for i := 0; i < srcPtr.NumField(); i++ {
		field := srcPtr.Type().Field(i)                  //获取到字段
		dic[field.Name] = srcPtr.FieldByName(field.Name) //将字段保存
	}

	for i := 0; i < destPtr.NumField(); i++ {
		currentField := destPtr.Type().Field(i)
		name := currentField.Name
		//如果与src中字段名匹配并且类型相同则赋值
		if dic[name].IsValid() && dic[name].Kind() == currentField.Type.Kind() && dic[name].CanSet() {
			destPtr.FieldByName(name).Set(dic[name])
		}
	}
	return dest
}

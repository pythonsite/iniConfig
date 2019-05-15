package iniConfig

import (
	"errors"
	"reflect"
	"fmt"
	"strconv"
	"strings"
)

// marshal 将data数据序列化为字节数组, data必须为struct
/*
func marshal(data interface{})(result []byte, err error) {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}
	var config []string
	valueInfo := reflect.ValueOf(data)
	for i:=0;i<typeInfo.NumField();i++ {
		sectionField := typeInfo.Field(i)
		sectionVal := valueInfo.Field(i)
		fieldType := sectionField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) ==0 {
			tagVal = sectionField.Name
		}
		var section string
		if i != 0 {
			section = fmt.Sprintf("\n[%s]\n", tagVal)
		} else {
			section = fmt.Sprintf("[%s]\n", tagVal)
		}
		config = append(config, section)
		for j :=0;j< fieldType.NumField();j++ {
			keyField := fieldType.Field(j)
			filedTagVal := keyField.Tag.Get("ini")
			if len(filedTagVal) == 0 {
				filedTagVal = keyField.Name
			}
			valueField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v")


		}

	}
}
 */

// unmarshal  fileData 表示获取的文件字节数组 confStruct 必须为struct 的ptr
func unmarshal(fileData []byte, confStruct interface{})(err error) {
	typeInfo := reflect.TypeOf(confStruct)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass ptr address")
		return
	}
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}
	lineArr := strings.Split(string(fileData),"\n")
	var sectionName string
	for lineNo, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		// ;和# 表示是注释
		if line[0] == ';' || line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			sectionName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v lineNo:%d",err,lineNo+1)
				return
			}
			continue
		}
		// 处理item
		err = parseItem(sectionName, line, confStruct)
		if err != nil {
			err = fmt.Errorf("%v lineNo:%d", err, lineNo + 1)
			return
		}
	}
	return
}

// parseSection用于处理每个section，获取[mysql] 中的值mysql
func parseSection(line string, typeInfo reflect.Type) (sectionName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("error invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("error invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionNameStr := strings.TrimSpace(line[1:len(line)-1])
		if len(sectionNameStr) == 0 {
			err = fmt.Errorf("error invalid sectionName:%s", line)
			return
		}
		for i:=0;i<typeInfo.NumField();i++ {
			field := typeInfo.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionNameStr {
				sectionName = field.Name
				fmt.Println("sectionName name:", sectionName)
				break
			}
		}
	}
	return
}

// parseItem 用于处理每个section下的配置信息
func parseItem(sectionName string, line string, confStruct interface{})(err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("sytax error not found '=' line:%s", line)
		return
	}
	key := strings.TrimSpace(line[0:index])
	value := strings.TrimSpace(line[index+1:])
	if len(key) == 0 {
		err = fmt.Errorf("sytax error, line:%s", line)
		return 
	}
	resultValue := reflect.ValueOf(confStruct)
	sectionValue := resultValue.Elem().FieldByName(sectionName)
	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", sectionName)
		return 
	}
	keyFieldName := ""
	for i:=0;i< sectionType.NumField();i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) ==0 {
		return 
	}
	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) {
		return 
	}
	switch fieldValue.Type().Kind(){
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		intVal,err2 := strconv.ParseInt(value,10, 64)
		if err2 != nil {
			err = err2
			return
		}
		fieldValue.SetInt(intVal)
	case reflect.Uint, reflect.Uint8,reflect.Uint16,reflect.Uint32, reflect.Uint64:
		uintVal, err2 := strconv.ParseUint(value, 10, 64)
		if err2 != nil {
			err = err2
			return
		}
		fieldValue.SetUint(uintVal)
	case reflect.Float32, reflect.Float64:
		floadVal, err2 := strconv.ParseFloat(value, 64)
		if err2 != nil {
			err = err2
			return
		}
		fieldValue.SetFloat(floadVal)
	default:
		err = fmt.Errorf("unsupport type:%v",fieldValue.Type().Kind())
		return
	}
	return
}




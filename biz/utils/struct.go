package utils

import (
	"encoding/json"
	"errors"
	"reflect"
)

func MapToStructByJson(m map[string]interface{}, s interface{}) error {
	targetValue := reflect.ValueOf(s)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, s)
	return nil
}

func SerializeStruct(s interface{}) (string, error) {
	targetValue := reflect.ValueOf(s)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return "", errors.New("target must be a non-nil pointer")
	}
	// 使用json.Marshal将response对象序列化为字节切片
	respBytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(respBytes), nil
}

func UnSerializeStruct(jsonStr string, s interface{}) error {
	jsonBytes := []byte(jsonStr)
	targetValue := reflect.ValueOf(s)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}
	if err := json.Unmarshal(jsonBytes, s); err != nil {
		return err
	}

	return nil
}

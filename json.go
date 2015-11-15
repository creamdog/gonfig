package gonfig

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"path"
	"strings"
)

type JsonGonfig struct {
	obj map[string]interface{}
}

func FromJson(reader io.Reader) (Gonfig, error) {
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &obj); err != nil {
		return nil, err
	}
	return &JsonGonfig{obj}, nil
}

func (jgonfig *JsonGonfig) GetString(key string, defaultValue interface{}) (string, error) {
	configValue, err := jgonfig.Get(key, defaultValue)
	if err != nil {
		return "", err
	}
	if stringValue, ok := configValue.(string); ok {
		return stringValue, nil
	} else {
		return "", &UnexpectedValueTypeError{key: key, value: configValue, message: "value is not a string"}
	}
}

func (jgonfig *JsonGonfig) GetInt(key string, defaultValue interface{}) (int, error) {
	value, err := jgonfig.GetFloat(key, defaultValue)
	if err != nil {
		return -1, err
	}
	return int(value), nil
}

func (jgonfig *JsonGonfig) GetFloat(key string, defaultValue interface{}) (float64, error) {
	configValue, err := jgonfig.Get(key, defaultValue)
	if err != nil {
		return -1.0, err
	}
	if floatValue, ok := configValue.(float64); ok {
		return floatValue, nil
	} else if intValue, ok := configValue.(int); ok {
		return float64(intValue), nil
	} else {
		return -1.0, &UnexpectedValueTypeError{key: key, value: configValue, message: "value is not a float"}
	}
}

func (jgonfig *JsonGonfig) GetBool(key string, defaultValue interface{}) (bool, error) {
	configValue, err := jgonfig.Get(key, defaultValue)
	if err != nil {
		return false, err
	}
	if boolValue, ok := configValue.(bool); ok {
		return boolValue, nil
	} else {
		return false, &UnexpectedValueTypeError{key: key, value: configValue, message: "value is not a bool"}
	}
}

func (jgonfig *JsonGonfig) GetAs(key string, target interface{}) error {
	configValue, err := jgonfig.Get(key, nil)
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(configValue)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return err
	}
	return nil
}

func (jgonfig *JsonGonfig) Get(key string, defaultValue interface{}) (interface{}, error) {
	parts := strings.Split(key, "/")
	var tmp interface{} = jgonfig.obj
	for index, part := range parts {
		if len(part) == 0 {
			continue
		}
		if confMap, ok := tmp.(map[string]interface{}); ok {
			if value, exists := confMap[part]; exists {
				tmp = value
			} else if defaultValue != nil {
				return defaultValue, nil
			} else {
				return nil, &KeyNotFoundError{key: path.Join(append(parts[:index], part)...)}
			}
		} else {
			return nil, &UnexpectedValueTypeError{key: path.Join(parts[:index]...), value: tmp, message: "value behind key is not a map[string]interface{}"}
		}
	}
	return tmp, nil
}

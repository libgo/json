package jsonx

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// func alias, add if needed.
var (
	Marshal             = json.Marshal
	MarshalIndent       = json.MarshalIndent
	MarshalToString     = json.MarshalToString
	Unmarshal           = json.Unmarshal
	UnmarshalFromString = json.UnmarshalFromString
	NewEncoder          = json.NewEncoder
	NewDecoder          = json.NewDecoder
	Valid               = json.Valid
	Get                 = json.Get
)

// MustMarshal
func MustMarshal(v interface{}) []byte {
	data, _ := Marshal(v)
	return data
}

// ToMap convert struct/bytes/string to map[string]interface{}
func ToMap(v interface{}) map[string]interface{} {
	if vm, ok := v.(map[string]interface{}); ok {
		return vm
	}

	buf := &bytes.Buffer{}
	switch t := v.(type) {
	case []byte:
		buf.Write(t)
	case string:
		buf.WriteString(t)
	default:
		err := NewEncoder(buf).Encode(v)
		if err != nil {
			return nil
		}
	}

	var result map[string]interface{}
	err := NewDecoder(buf).Decode(&result)
	if err != nil {
		return nil
	}

	return result
}

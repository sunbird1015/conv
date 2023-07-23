package conv

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Conv(v interface{}, o interface{}) error {
	config := &DecoderConfig{
		Result:           o,
		WeaklyTypedInput: true,
		Squash:           true,
	}
	decoder, err := NewDecoder(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = decoder.decode("", v, reflect.ValueOf(decoder.config.Result).Elem())
	if err != nil {
		//fmt.Println(err.Error())
		return err
	}
	return nil
}

func Int(v interface{}) int {
	var output int
	Conv(v, &output)
	return output
}

func Ints(v interface{}) []int {
	var output []int
	Conv(v, &output)
	return output
}

func Int8(v interface{}) int8 {
	var output int8
	Conv(v, &output)
	return output
}

func Int8s(v interface{}) []int8 {
	var output []int8
	Conv(v, &output)
	return output
}

func Int16(v interface{}) int16 {
	var output int16
	Conv(v, &output)
	return output
}

func Int16s(v interface{}) []int16 {
	var output []int16
	Conv(v, &output)
	return output
}

func Int32(v interface{}) int32 {
	var output int32
	Conv(v, &output)
	return output
}

func Int32s(v interface{}) []int32 {
	var output []int32
	Conv(v, &output)
	return output
}

func Int64(v interface{}) int64 {
	var output int64
	Conv(v, &output)
	return output
}

func Int64s(v interface{}) []int64 {
	var output []int64
	Conv(v, &output)
	return output
}

func Uint(v interface{}) uint {
	var output uint
	Conv(v, &output)
	return output
}

func Uints(v interface{}) []uint {
	var output []uint
	Conv(v, &output)
	return output
}

func Uint8(v interface{}) uint8 {
	var output uint8
	Conv(v, &output)
	return output
}

func Uint8s(v interface{}) []uint8 {
	var output []uint8
	Conv(v, &output)
	return output
}

func Uint16(v interface{}) uint16 {
	var output uint16
	Conv(v, &output)
	return output
}

func Uint16s(v interface{}) []uint16 {
	var output []uint16
	Conv(v, &output)
	return output
}

func Uint32(v interface{}) uint32 {
	var output uint32
	Conv(v, &output)
	return output
}

func Uint32s(v interface{}) []uint32 {
	var output []uint32
	Conv(v, &output)
	return output
}

func Uint64(v interface{}) uint64 {
	var output uint64
	Conv(v, &output)
	return output
}

func Uint64s(v interface{}) []uint64 {
	var output []uint64
	Conv(v, &output)
	return output
}

func Bool(v interface{}) bool {
	var output bool
	Conv(v, &output)
	return output
}

func Bools(v interface{}) []bool {
	var output []bool
	Conv(v, &output)
	return output
}

func String(v interface{}) string {
	var output string
	if err := Conv(v, &output); err != nil {
		if ret, err := json.Marshal(v); err == nil {
			return string(ret)
		}
	}
	return output
}

func Strings(v interface{}) []string {
	var output []string
	Conv(v, &output)
	return output
}

func Map(v interface{}) map[string]interface{} {
	var output map[string]interface{}
	Conv(v, &output)
	return output
}

func MapStrStr(v interface{}) map[string]string {
	var output map[string]string
	Conv(v, &output)
	return output
}

func Maps(v interface{}) []map[string]interface{} {
	var output []map[string]interface{}
	Conv(v, &output)
	return output
}

func Interfaces(v interface{}) []interface{} {
	var output []interface{}
	Conv(v, &output)
	return output
}

func Struct(v interface{}, o interface{}) {
	Conv(v, o)
}

func Float32(v interface{}) float32 {
	var output float32
	Conv(v, &output)
	return output
}

func Float32s(v interface{}) []float32 {
	var output []float32
	Conv(v, &output)
	return output
}

func Float64(v interface{}) float64 {
	var output float64
	Conv(v, &output)
	return output
}

func Float64s(v interface{}) []float64 {
	var output []float64
	Conv(v, &output)
	return output
}

func Object(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	data, err := json.Marshal(v)
	if err != nil {
		return v
	}
	var json_doc interface{}
	err = json.Unmarshal(data, &json_doc)
	if err != nil {
		return v
	}
	return json_doc
}

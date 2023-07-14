package conv

import (
	"reflect"
)

func conv(v interface{}, o interface{}) error {
	config := &DecoderConfig{
		Result:           o,
		WeaklyTypedInput: true,
	}
	decoder, err := NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.decode("", v, reflect.ValueOf(decoder.config.Result).Elem())
	if err != nil {
		return err
	}
	return nil
}

func Int(v interface{}) (int, error) {
	var output int
	err := conv(v, &output)
	return output, err
}

func Ints(v interface{}) ([]int, error) {
	var output []int
	err := conv(v, &output)
	return output, err
}

func Int8(v interface{}) (int8, error) {
	var output int8
	err := conv(v, &output)
	return output, err
}

func Int8s(v interface{}) ([]int8, error) {
	var output []int8
	err := conv(v, &output)
	return output, err
}

func Int16(v interface{}) (int16, error) {
	var output int16
	err := conv(v, &output)
	return output, err
}

func Int16s(v interface{}) ([]int16, error) {
	var output []int16
	err := conv(v, &output)
	return output, err
}

func Int32(v interface{}) (int32, error) {
	var output int32
	err := conv(v, &output)
	return output, err
}

func Int32s(v interface{}) ([]int32, error) {
	var output []int32
	err := conv(v, &output)
	return output, err
}

func Int64(v interface{}) (int64, error) {
	var output int64
	err := conv(v, &output)
	return output, err
}

func Int64s(v interface{}) ([]int64, error) {
	var output []int64
	err := conv(v, &output)
	return output, err
}

func Uint(v interface{}) (uint, error) {
	var output uint
	err := conv(v, &output)
	return output, err
}

func Uints(v interface{}) ([]uint, error) {
	var output []uint
	err := conv(v, &output)
	return output, err
}

func Uint8(v interface{}) (uint8, error) {
	var output uint8
	err := conv(v, &output)
	return output, err
}

func Uint8s(v interface{}) ([]uint8, error) {
	var output []uint8
	err := conv(v, &output)
	return output, err
}

func Uint16(v interface{}) (uint16, error) {
	var output uint16
	err := conv(v, &output)
	return output, err
}

func Uint16s(v interface{}) ([]uint16, error) {
	var output []uint16
	err := conv(v, &output)
	return output, err
}

func Uint32(v interface{}) (uint32, error) {
	var output uint32
	err := conv(v, &output)
	return output, err
}

func Uint32s(v interface{}) ([]uint32, error) {
	var output []uint32
	err := conv(v, &output)
	return output, err
}

func Uint64(v interface{}) (uint64, error) {
	var output uint64
	err := conv(v, &output)
	return output, err
}

func Uint64s(v interface{}) ([]uint64, error) {
	var output []uint64
	err := conv(v, &output)
	return output, err
}

func Bool(v interface{}) (bool, error) {
	var output bool
	err := conv(v, &output)
	return output, err
}

func Bools(v interface{}) ([]bool, error) {
	var output []bool
	err := conv(v, &output)
	return output, err
}

func String(v interface{}) (string, error) {
	var output string
	err := conv(v, &output)
	return output, err
}

func Strings(v interface{}) ([]string, error) {
	var output []string
	err := conv(v, &output)
	return output, err
}

func Map(v interface{}) (map[string]interface{}, error) {
	var output map[string]interface{}
	err := conv(v, &output)
	return output, err
}

func MapStrStr(v interface{}) (map[string]string, error) {
	var output map[string]string
	err := conv(v, &output)
	return output, err
}

func Maps(v interface{}) ([]map[string]interface{}, error) {
	var output []map[string]interface{}
	err := conv(v, &output)
	return output, err
}

func Interfaces(v interface{}) ([]interface{}, error) {
	var output []interface{}
	err := conv(v, &output)
	return output, err
}

func Struct(v interface{}, o interface{}) error {
	return conv(v, o)
}

func Float32(v interface{}) (float32, error) {
	var output float32
	err := conv(v, &output)
	return output, err
}

func Float32s(v interface{}) ([]float32, error) {
	var output []float32
	err := conv(v, &output)
	return output, err
}

func Float64(v interface{}) (float64, error) {
	var output float64
	err := conv(v, &output)
	return output, err
}

func Float64s(v interface{}) ([]float64, error) {
	var output []float64
	err := conv(v, &output)
	return output, err
}

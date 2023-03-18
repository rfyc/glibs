package conv

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type stringInterface interface {
	String() string
}

func Format(temp interface{}, objType string) interface{} {

	switch strings.ToLower(objType) {
	case "int":
		return Int(temp)
	case "int8":
		return Int8(temp)
	case "int16":
		return Int16(temp)
	case "int32":
		return Int32(temp)
	case "int64":
		return Int64(temp)
	case "uint":
		return Uint(temp)
	case "uint8":
		return Uint8(temp)
	case "uint16":
		return Uint16(temp)
	case "uint32":
		return Uint32(temp)
	case "uint64":
		return Uint64(temp)
	case "float32":
		return Float32(temp)
	case "float64":
		return Float64(temp)
	case "bool":
		return Bool(temp)
	case "string":
		return String(temp)
	case "[]bype":
		return Bytes(temp)
	default:
		return temp
	}
}

func String(temp interface{}) string {

	if temp == nil {
		return ""
	}
	switch value := temp.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.Itoa(int(value))
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		return value.Format("2006-01-02 15:04:05")
	default:
		if obj, ok := value.(stringInterface); ok {
			return obj.String()
		} else {
			// 默认使用json进行字符串转换
			result, _ := json.Marshal(value)
			return string(result)
		}
	}
}

func btoi(b bool) uint8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func Int(temp interface{}) int {
	if b, ok := temp.(bool); ok {
		return int(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseInt(result, 10, 64)
	return int(number)
}
func Int8(temp interface{}) int8 {
	if b, ok := temp.(bool); ok {
		return int8(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseInt(result, 10, 8)
	return int8(number)
}
func Int16(temp interface{}) int16 {
	if b, ok := temp.(bool); ok {
		return int16(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseInt(result, 10, 16)
	return int16(number)
}
func Int32(temp interface{}) int32 {
	if b, ok := temp.(bool); ok {
		return int32(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseInt(result, 10, 32)
	return int32(number)
}
func Int64(temp interface{}) int64 {
	if b, ok := temp.(bool); ok {
		return int64(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseInt(result, 10, 64)
	return int64(number)
}
func Uint(temp interface{}) uint {
	if b, ok := temp.(bool); ok {
		return uint(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseUint(result, 10, 64)
	return uint(number)
}
func Uint8(temp interface{}) uint8 {
	if b, ok := temp.(bool); ok {
		return uint8(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseUint(result, 10, 8)
	return uint8(number)
}
func Uint16(temp interface{}) uint16 {
	if b, ok := temp.(bool); ok {
		return uint16(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseUint(result, 10, 16)
	return uint16(number)
}
func Uint32(temp interface{}) uint32 {
	if b, ok := temp.(bool); ok {
		return uint32(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseUint(result, 10, 32)
	return uint32(number)
}
func Uint64(temp interface{}) uint64 {
	if b, ok := temp.(bool); ok {
		return uint64(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseUint(result, 10, 64)
	return uint64(number)
}
func Float32(temp interface{}) float32 {
	if b, ok := temp.(bool); ok {
		return float32(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseFloat(result, 32)
	return float32(number)
}
func Float64(temp interface{}) float64 {
	if b, ok := temp.(bool); ok {
		return float64(btoi(b))
	}
	result := String(temp)
	number, _ := strconv.ParseFloat(result, 64)
	return float64(number)
}
func Bool(temp interface{}) bool {
	result, _ := strconv.ParseBool(String(temp))
	return result
}
func Bytes(temp interface{}) []byte {
	return []byte(String(temp))
}

func Strings(temp interface{}) []string {
	var ss []string
	json.Unmarshal(Bytes(temp), &ss)
	return ss
}

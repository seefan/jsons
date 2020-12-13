package jsons

import "strconv"

func escape(json []byte) []byte {
	size := len(json)
	var str []byte
	for i := 0; i < size; i++ {
		switch json[i] {
		case '\\':
			str = append(str, '\\', '\\')
		case '/':
			str = append(str, '\\', '/')
		case 'b':
			str = append(str, '\\', 'b')
		case 'f':
			str = append(str, '\\', 'f')
		case 'n':
			str = append(str, '\\', 'n')
		case 'r':
			str = append(str, '\\', 'r')
		case 't':
			str = append(str, '\\', 't')
		case '"':
			str = append(str, '\\', '"')
		default:
			str = append(str, json[i])
		}
	}
	return str
}
func buildValue(value interface{}) JsonByte {
	jv := &simpleValue{}
	switch value := value.(type) {
	case string:
		jv.val = append([]byte{byte('"')}, escape([]byte(value))...)
		jv.val = append(jv.val, byte('"'))
	case []byte:
		jv.val = append([]byte{byte('"')}, escape(value)...)
		jv.val = append(jv.val, byte('"'))
	case int:
		jv.val = strconv.AppendInt(nil, int64(value), 10)
	case int8:
		jv.val = []byte{byte(value)}
	case int16:
		jv.val = strconv.AppendInt(nil, int64(value), 10)
	case int32:
		jv.val = strconv.AppendInt(nil, int64(value), 10)
	case int64:
		jv.val = strconv.AppendInt(nil, value, 10)
	case uint8:
		jv.val = []byte{byte(value)}
	case uint16:
		jv.val = strconv.AppendUint(nil, uint64(value), 10)
	case uint32:
		jv.val = strconv.AppendUint(nil, uint64(value), 10)
	case uint64:
		jv.val = strconv.AppendUint(nil, value, 10)
	case float32:
		jv.val = strconv.AppendFloat(nil, float64(value), 'g', -1, 32)
	case float64:
		jv.val = strconv.AppendFloat(nil, value, 'g', -1, 64)
	case bool:
		if value {
			jv.val = []byte("true")
		} else {
			jv.val = []byte("false")
		}
	case nil:
		jv.val = []byte("null")
	default:
		panic("unsupported type")
	}
	return jv
}

package jsons

type JsonWriteArray struct {
	val []JsonByte
}

func (j *JsonWriteArray) Bytes() []byte {
	if j.val == nil {
		return nil
	}
	var s []byte = []byte("[")
	has := false
	for _, v := range j.val {
		if !has {
			has = true
		} else {
			s = append(s, byte(','))
		}
		s = append(s, v.Bytes()...)
	}
	s = append(s, byte(']'))
	return s
}

func NewJsonArray() *JsonWriteArray {
	return &JsonWriteArray{}
}

func (j *JsonWriteArray) Add(value ...interface{}) *JsonWriteArray {
	for _, v := range value {
		switch v := v.(type) {
		case JsonByte:
			j.val = append(j.val, v)
		default:
			jv := buildValue(v)
			j.val = append(j.val, jv)
		}
	}
	return j
}

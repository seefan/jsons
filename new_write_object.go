package jsons

type JsonByte interface {
	Bytes() []byte
}
type simpleValue struct {
	val []byte
}

func (t *simpleValue) Bytes() []byte {
	return t.val
}

type JsonWriteObject struct {
	val map[string]JsonByte
}

func (j *JsonWriteObject) Bytes() []byte {
	if j.val == nil {
		return nil
	}
	var s []byte = []byte("{")
	has := false
	for k, v := range j.val {
		if !has {
			has = true
			s = append(s, byte('"'))
		} else {
			s = append(s, []byte(`,"`)...)
		}

		s = append(s, []byte(k)...)
		s = append(s, byte('"'), byte(':'))
		s = append(s, v.Bytes()...)
	}
	s = append(s, byte('}'))
	return s
}
func (j *JsonWriteObject) Put(key string, value interface{}) *JsonWriteObject {
	switch value := value.(type) {
	case JsonByte:
		j.val[key] = value
	default:
		j.val[key] = buildValue(value)
	}
	return j
}

func NewJsonObject() *JsonWriteObject {
	return &JsonWriteObject{val: make(map[string]JsonByte)}
}

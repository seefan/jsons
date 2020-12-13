/*
@Time : 2019-03-27 13:54
@Author : seefan
@File : JsonObject
@Software: jsons
*/
package jsons

type JsonReadObject struct {
	reader
	val map[string]JsonValue
	obj map[string]*JsonReadObject
	arr map[string]*JsonReadArray
}

func (j *JsonReadObject) parse() {
	if !j.validObject() {
		j.end = -1
		return
	}

	//remove {}
	j.index++
	j.end--
	for j.index <= j.end {
		j.skip()
		if !j.IsValid() {
			break
		}
		key := string(unescape(j.parseString()))
		j.skipSplit()
		if !j.IsValid() {
			break
		}
		value := j.parseValue()
		j.val[key] = JsonValue(value)
		if !j.hasMore() {
			break
		}
	}
}

func ParseJsonObject(bs []byte) *JsonReadObject {
	j := &JsonReadObject{
		reader: *newReader(bs),
		val:    make(map[string]JsonValue),
	}
	j.parse()
	return j
}
func (j *JsonReadObject) C(name string) bool {
	return j.Contains(name)
}

//Keys get all key
func (j *JsonReadObject) Keys() []string {
	var keys []string
	for k := range j.val {
		keys = append(keys, k)
	}
	for k := range j.obj {
		keys = append(keys, k)
	}
	for k := range j.arr {
		keys = append(keys, k)
	}
	return keys
}
func (j *JsonReadObject) Contains(name string) bool {
	if _, ok := j.val[name]; ok {
		return true
	}
	return false
}
func (j *JsonReadObject) Left(name string, size int) string {
	if j.V(name).IsEmpty() {
		return ""
	}
	bs := j.GetValue(name).Bytes()
	r := []rune(string(bs))
	if len(r) < size {
		return string(bs)
	} else {
		return string(r[0:size])
	}
}
func (j *JsonReadObject) V(name string) JsonValue {
	return j.GetValue(name)
}
func (j *JsonReadObject) GetValue(name string) JsonValue {
	return j.val[name]
}
func (j *JsonReadObject) O(name string) *JsonReadObject {
	return j.GetObject(name)
}
func (j *JsonReadObject) GetObject(name string) *JsonReadObject {
	if j.obj != nil {
		if v, ok := j.obj[name]; ok {
			return v
		}
	}
	return j.GetObjectForce(name)
}
func (j *JsonReadObject) GetObjectForce(name string) *JsonReadObject {
	if v, ok := j.val[name]; ok {
		if j.obj == nil {
			j.obj = make(map[string]*JsonReadObject)
		}
		j.obj[name] = ParseJsonObject(v.Bytes())
		return j.obj[name]
	}
	return &JsonReadObject{
		val: make(map[string]JsonValue),
	}
}
func (j *JsonReadObject) A(name string) *JsonReadArray {
	return j.GetArray(name)
}
func (j *JsonReadObject) GetArray(name string) *JsonReadArray {
	if j.arr != nil {
		if v, ok := j.arr[name]; ok {
			return v
		}
	}
	return j.GetArrayForce(name)
}
func (j *JsonReadObject) GetArrayForce(name string) *JsonReadArray {
	if v, ok := j.val[name]; ok {
		if j.arr == nil {
			j.arr = make(map[string]*JsonReadArray)
		}
		j.arr[name] = ParseJsonArray(v.Bytes())
		return j.arr[name]
	}
	return &JsonReadArray{}
}

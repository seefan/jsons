/*
@Time : 2019-03-27 18:27
@Author : seefan
@File : jsonarr
@Software: jsons
*/
package jsons

import (
	"bytes"
)

type JsonReadArray struct {
	reader
	arr []JsonValue
}

func (j *JsonReadArray) Get(i int) JsonValue {
	return j.arr[i]
}
func (j *JsonReadArray) Size() int {
	return len(j.arr)
}
func ParseJsonArray(bs []byte) *JsonReadArray {
	j := &JsonReadArray{
		reader: *newReader(bs),
	}
	j.parse()
	return j
}
func (j *JsonReadArray) Each(f func(int, JsonValue)) {
	if j.arr != nil {
		for i, v := range j.arr {
			f(i, v)
		}
	}
}
func (j *JsonReadArray) parse() {
	if !j.validArray() {
		j.LastError = "JsonArray format error"
		j.end = -1
		return
	}
	//remove []
	//可以考虑使用，可能性能会更高strings.Index()
	j.index++
	j.end--

	j.skip()
	start := j.index
	str := 0
	depth := 0
	for j.index < j.end {
		switch j.data[j.index] {
		case '[', '{':
			if str%2 == 0 {
				depth++
			}
		case ']', '}':
			if str%2 == 0 {
				depth--
			}
		case '"':
			str++
		case '\\':
			if j.index+1 < j.end && j.data[j.index+1] == '"' {
				j.index++
			}
		case ',':
			if depth == 0 && str%2 == 0 {
				if j.data[start] == '"' {
					j.arr = append(j.arr, JsonValue(j.data[start+1:j.index-1]))
				} else {
					j.arr = append(j.arr, JsonValue(bytes.TrimSpace(j.data[start:j.index])))
				}
				start = j.index + 1
			}
		}
		j.index++
	}
	if start < j.index {
		if j.data[start] == '"' {
			j.arr = append(j.arr, JsonValue(j.data[start+1:j.index-1]))
		} else {
			j.arr = append(j.arr, JsonValue(j.data[start:j.index]))
		}
	}
}

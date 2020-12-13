package jsons_test

import (
	"testing"

	"github.com/seefan/jsons"
)

func TestWrite(t *testing.T) {
	js := jsons.NewJsonObject().Put("a", 1).Put("b", "2")
	bs := js.Bytes()
	t.Log(bs)
}

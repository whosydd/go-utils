package goutils

import (
	"reflect"
	"testing"
)

func TestIsExist(t *testing.T) {
	dst := "./examples"
	got, err := IsExist(dst)
	if err != nil {
		t.Error(err)
	}
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v is not exist. want:%v got:%v", dst, want, got)
	}
}

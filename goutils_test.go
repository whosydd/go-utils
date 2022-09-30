package goutils

import (
	"reflect"
	"testing"
)

func TestIsExist(t *testing.T) {
	test := []string{
		"./examples",
		"./go.mod",
		// "./test.txt",
		"d:/desktop",
		// "d:/desktop/test",
		// "d:/desktop/test.txt",
	}

	for _, dst := range test {
		t.Run(dst, func(t *testing.T) {
			got, err := IsExist(dst)
			if err != nil {
				t.Error(err)
			}
			want := true
			if !reflect.DeepEqual(got, want) {
				t.Errorf("%v is not exist. want:%v got:%v", dst, want, got)
			}
		})
	}
}

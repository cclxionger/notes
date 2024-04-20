package split

import (
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v got: %v", want, got)
	}
}

package findMax

import (
	"reflect"
	"testing"
)

func TestFindMax(t *testing.T) {
	type args struct {
		comp     func(el1, el2 interface{}) bool
		elements []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"Int",
			args{func(el1, el2 interface{}) bool {
				return el1.(int) > el2.(int)
			}, []interface{}{1, 2, 3, 4, 100}},
			100,
			false,
		}, {
			"String",
			args{func(el1, el2 interface{}) bool {
				return el1.(string) > el2.(string)
			}, []interface{}{"a", "b", "c"}},
			"c",
			false,
		}, {
			"IntSliceLenght",
			args{func(el1, el2 interface{}) bool {
				return len(el1.([]int)) > len(el2.([]int))
			}, []interface{}{[]int{1}, []int{2}, []int{3}, []int{4}, []int{1, 100}}},
			[]int{1, 100},
			false,
		}, {
			"Float",
			args{func(el1, el2 interface{}) bool {
				return el1.(float64) > el2.(float64)
			}, []interface{}{1.0, 2.0, 3.0, 4.0, 100.0}},
			100.0,
			false,
		}, {
			"Int is not a string Error",
			args{func(el1, el2 interface{}) bool {
				return el1.(string) > el2.(string)
			}, []interface{}{1, 2, 3, 4, 100}},
			nil,
			true,
		}, {
			"Not enough parameters Error",
			args{func(el1, el2 interface{}) bool {
				return el1.(string) > el2.(string)
			}, []interface{}{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMax(tt.args.comp, tt.args.elements)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

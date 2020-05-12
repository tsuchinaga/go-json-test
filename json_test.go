package go_json_test

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	type target struct {
		Str string `json:"str,omitempty"`
	}

	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		{desc: "空文字", target: &target{Str: ""}, expect: []byte("{}")},
		{desc: "nullという文字列", target: &target{Str: "null"}, expect: []byte(`{"str":"null"}`)},
		{desc: "foo", target: &target{Str: "foo"}, expect: []byte(`{"str":"foo"}`)},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	type target struct {
		Int int `json:"int,omitempty"`
	}

	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		{desc: "-1", target: &target{Int: -1}, expect: []byte(`{"int":-1}`)},
		{desc: "0", target: &target{Int: 0}, expect: []byte(`{}`)},
		{desc: "1", target: &target{Int: 1}, expect: []byte(`{"int":1}`)},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestFloat(t *testing.T) {
	t.Parallel()

	type target struct {
		Float float64 `json:"float,omitempty"`
	}

	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		{desc: "-0.5", target: &target{Float: -0.5}, expect: []byte(`{"float":-0.5}`)},
		{desc: "0", target: &target{Float: 0}, expect: []byte(`{}`)},
		{desc: "0.5", target: &target{Float: 0.5}, expect: []byte(`{"float":0.5}`)},
		// {desc: "Infinity", target: &target{Float: math.Inf(0)}, expect: []byte(`{"float":"+Inf"}`)}, // Infinityは対応してない
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestBool(t *testing.T) {
	t.Parallel()

	type target struct {
		Bool bool `json:"bool,omitempty"`
	}

	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		{desc: "true", target: &target{Bool: true}, expect: []byte(`{"bool":true}`)},
		{desc: "false", target: &target{Bool: false}, expect: []byte(`{}`)}, // falseはomitemptyで消される
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	t.Parallel()

	type data struct {
		Num int `json:"num,omitempty"`
	}

	type target struct {
		Slice []*data `json:"slice,omitempty"`
	}

	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		{desc: "nil", target: &target{Slice: nil}, expect: []byte(`{}`)},
		{desc: "空", target: &target{Slice: []*data{}}, expect: []byte(`{}`)},
		{desc: "構成要素がnil", target: &target{Slice: []*data{nil}}, expect: []byte(`{"slice":[null]}`)},
		{desc: "構成要素が有効な値", target: &target{Slice: []*data{{Num: 1}}}, expect: []byte(`{"slice":[{"num":1}]}`)},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	type data struct {
		Num int `json:"num,omitempty"`
	}

	type target struct {
		Map map[string]*data `json:"map,omitempty"`
	}

	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		{desc: "nil", target: &target{Map: nil}, expect: []byte(`{}`)},
		{desc: "空", target: &target{Map: map[string]*data{}}, expect: []byte(`{}`)},
		{desc: "keyが空文字", target: &target{Map: map[string]*data{"": {Num: 1}}}, expect: []byte(`{"map":{"":{"num":1}}}`)},
		{desc: "valueがnil", target: &target{Map: map[string]*data{"key": nil}}, expect: []byte(`{"map":{"key":null}}`)},
		{desc: "有効なデータ", target: &target{Map: map[string]*data{"key": {Num: -1}}}, expect: []byte(`{"map":{"key":{"num":-1}}}`)},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestInterface(t *testing.T) {
	t.Parallel()

	type target struct {
		Interface interface{} `json:"interface,omitempty"`
	}

	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		{desc: "nil", target: &target{Interface: nil}, expect: []byte(`{}`)},
		{desc: "string,omitempty", target: &target{Interface: ""}, expect: []byte(`{"interface":""}`)},
		{desc: "string,omitempty", target: &target{Interface: "string"}, expect: []byte(`{"interface":"string"}`)},
		{desc: "int,omitempty", target: &target{Interface: 0}, expect: []byte(`{"interface":0}`)},
		{desc: "int", target: &target{Interface: 1}, expect: []byte(`{"interface":1}`)},
		{desc: "float,omitempty", target: &target{Interface: 0.0}, expect: []byte(`{"interface":0}`)},
		{desc: "float", target: &target{Interface: 0.5}, expect: []byte(`{"interface":0.5}`)},
		{desc: "bool,omitempty", target: &target{Interface: false}, expect: []byte(`{"interface":false}`)},
		{desc: "bool", target: &target{Interface: true}, expect: []byte(`{"interface":true}`)},
		{desc: "slice,omitempty", target: &target{Interface: []int{}}, expect: []byte(`{"interface":[]}`)},
		{desc: "slice", target: &target{Interface: []int{1, 2, 3}}, expect: []byte(`{"interface":[1,2,3]}`)},
		{desc: "map,omitempty", target: &target{Interface: map[string]int{}}, expect: []byte(`{"interface":{}}`)},
		{desc: "map", target: &target{Interface: map[string]int{"foo": 1, "bar": 2, "baz": 3}}, expect: []byte(`{"interface":{"bar":2,"baz":3,"foo":1}}`)},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestFunc(t *testing.T) {
	t.Parallel()

	type target struct {
		Func interface{} `json:"func,omitempty"`
	}

	// funcは対応していない
	testTable := []struct {
		desc   string
		target *target
		expect []byte
	}{
		// {desc: "func()", target: &target{Func: func() {}}, expect: []byte(`{}`)},
		// {desc: "func()string", target: &target{Func: func() string { return "foo" }}, expect: []byte(`{}`)},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual, err := json.Marshal(test.target)
			if err != nil {
				t.Errorf("%s エラー\n%v\n", t.Name(), err)
			}
			if !reflect.DeepEqual(test.expect, actual) {
				t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), test.expect, actual)
			}
		})
	}
}

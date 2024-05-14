package xmap

import (
	"encoding/json"
	"testing"
)

var jsonStr = `
{
	"a":"b",
	"int":1,
	"float":1.1,
	"float2":1.8,
	"k": {
		"k2": "v",
		"k3.k4": -4
	}
}
`
var data map[string]any

func init() {
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}
}

func Test_Result(t *testing.T) {

	v := GetMapValue(data, "a")
	if !v.Exists() {
		t.Errorf("Expected to find key 'a' in map")
	}
	if v.String() != "b" {
		t.Errorf("Expected value to be 'b', got %v", v.String())
	}

	v = GetMapValue(data, "int")
	if !v.Exists() {
		t.Errorf("Expected to find key 'int' in map")
	}
	if v.Int() != 1 {
		t.Errorf("Expected value to be 1, got %v", v.Int())
	}

	v = GetMapValue(data, "float")
	if !v.Exists() {
		t.Errorf("Expected to find key 'float' in map")
	}
	if v.Float() != 1.1 {
		t.Errorf("Expected value to be 1, got %v", v.Float())
	}
	if v.Int() != 1 {
		t.Errorf("Expected value to be 1, got %v", v.Int())
	}

	v = GetMapValue(data, "float2")
	if !v.Exists() {
		t.Errorf("Expected to find key 'float' in map")
	}
	if v.Float() != 1.8 {
		t.Errorf("Expected value to be 1, got %v", v.Float())
	}
	if v.Int() != 1 {
		t.Errorf("Expected value to be 2, got %v", v.Int())
	}

	v = GetMapValue(data, "float.notexist")
	if v.Exists() {
		t.Errorf("Expected to not find key 'float_notexist' in map")
	}

	v = GetMapValue(data, "k.k3.k4")
	// t.Log(v.raw)
	if !v.Exists() {
		t.Errorf("Expected to find key 'k.k3.k4' in map")
	}
	if v.String() != "-4" {
		t.Errorf("Expected value to be 'v4', got %v", v.String())
	}
}
